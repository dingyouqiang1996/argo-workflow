package common

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/argoproj/argo/errors"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/valyala/fasttemplate"
)

// wfValidationCtx is the context for validating a workflow spec
type wfValidationCtx struct {
	wf *wfv1.Workflow
	// globalParams keeps track of variables which are available the global
	// scope and can be referenced from anywhere.
	globalParams map[string]string
	// results tracks if validation has already been run on a template
	results map[string]bool
}

const (
	// placeholderValue is an arbitrary string to perform mock substitution of variables
	placeholderValue = "placeholder"

	// anyItemMagicValue is a magic value set in addItemsToScope() and checked in
	// resolveAllVariables() to determine if any {{item.name}} can be accepted during
	// variable resolution (to support withParam)
	anyItemMagicValue = "item.*"
)

// ValidateWorkflow accepts a workflow and performs validation against it
func ValidateWorkflow(wf *wfv1.Workflow) error {
	ctx := wfValidationCtx{
		wf:           wf,
		globalParams: make(map[string]string),
		results:      make(map[string]bool),
	}

	err := validateWorkflowFieldNames(wf.Spec.Templates)
	if err != nil {
		return errors.Errorf(errors.CodeBadRequest, "spec.templates%s", err.Error())
	}
	err = validateArguments("spec.arguments.", wf.Spec.Arguments)
	if err != nil {
		return err
	}
	ctx.globalParams[GlobalVarWorkflowName] = placeholderValue
	ctx.globalParams[GlobalVarWorkflowUID] = placeholderValue
	for _, param := range ctx.wf.Spec.Arguments.Parameters {
		ctx.globalParams["workflow.parameters."+param.Name] = placeholderValue
	}
	if ctx.wf.Spec.Entrypoint == "" {
		return errors.New(errors.CodeBadRequest, "spec.entrypoint is required")
	}
	entryTmpl := ctx.wf.GetTemplate(ctx.wf.Spec.Entrypoint)
	if entryTmpl == nil {
		return errors.Errorf(errors.CodeBadRequest, "spec.entrypoint template '%s' undefined", ctx.wf.Spec.Entrypoint)
	}
	err = ctx.validateTemplate(entryTmpl, ctx.wf.Spec.Arguments)
	if err != nil {
		return err
	}
	if ctx.wf.Spec.OnExit != "" {
		exitTmpl := ctx.wf.GetTemplate(ctx.wf.Spec.OnExit)
		if exitTmpl == nil {
			return errors.Errorf(errors.CodeBadRequest, "spec.onExit template '%s' undefined", ctx.wf.Spec.OnExit)
		}
		// now when validating onExit, {{workflow.status}} is now available as a global
		ctx.globalParams[GlobalVarWorkflowStatus] = placeholderValue
		err = ctx.validateTemplate(exitTmpl, ctx.wf.Spec.Arguments)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ctx *wfValidationCtx) validateTemplate(tmpl *wfv1.Template, args wfv1.Arguments) error {
	_, ok := ctx.results[tmpl.Name]
	if ok {
		// we already processed this template
		return nil
	}
	ctx.results[tmpl.Name] = true
	if err := validateTemplateType(tmpl); err != nil {
		return err
	}
	_, err := ProcessArgs(tmpl, args, ctx.globalParams, true)
	if err != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s %s", tmpl.Name, err)
	}
	scope, err := validateInputs(tmpl)
	if err != nil {
		return err
	}
	for globalVar, val := range ctx.globalParams {
		scope[globalVar] = val
	}
	switch tmpl.GetType() {
	case wfv1.TemplateTypeSteps:
		err = ctx.validateSteps(scope, tmpl)
	case wfv1.TemplateTypeDAG:
		err = ctx.validateDAG(scope, tmpl)
	default:
		err = validateLeaf(scope, tmpl)
	}
	if err != nil {
		return err
	}
	err = validateOutputs(scope, tmpl)
	if err != nil {
		return err
	}
	return nil
}

// validateTemplateType validates that only one template type is defined
func validateTemplateType(tmpl *wfv1.Template) error {
	numTypes := 0
	for _, tmplType := range []interface{}{tmpl.Container, tmpl.Steps, tmpl.Script, tmpl.Resource, tmpl.DAG} {
		if !reflect.ValueOf(tmplType).IsNil() {
			numTypes++
		}
	}
	switch numTypes {
	case 0:
		return errors.New(errors.CodeBadRequest, "template type unspecified. choose one of: container, steps, script, resource, dag")
	case 1:
	default:
		return errors.New(errors.CodeBadRequest, "multiple template types specified. choose one of: container, steps, script, resource, dag")
	}
	return nil
}

func validateInputs(tmpl *wfv1.Template) (map[string]interface{}, error) {
	err := validateWorkflowFieldNames(tmpl.Inputs.Parameters)
	if err != nil {
		return nil, errors.Errorf(errors.CodeBadRequest, "templates.%s.inputs.parameters%s", tmpl.Name, err.Error())
	}
	err = validateWorkflowFieldNames(tmpl.Inputs.Artifacts)
	if err != nil {
		return nil, errors.Errorf(errors.CodeBadRequest, "templates.%s.inputs.artifacts%s", tmpl.Name, err.Error())
	}
	scope := make(map[string]interface{})
	for _, param := range tmpl.Inputs.Parameters {
		scope[fmt.Sprintf("inputs.parameters.%s", param.Name)] = true
	}
	isLeaf := tmpl.Container != nil || tmpl.Script != nil
	for _, art := range tmpl.Inputs.Artifacts {
		artRef := fmt.Sprintf("inputs.artifacts.%s", art.Name)
		scope[artRef] = true
		if isLeaf {
			if art.Path == "" {
				return nil, errors.Errorf(errors.CodeBadRequest, "templates.%s.%s.path not specified", tmpl.Name, artRef)
			}
		} else {
			if art.Path != "" {
				return nil, errors.Errorf(errors.CodeBadRequest, "templates.%s.%s.path only valid in container/script templates", tmpl.Name, artRef)
			}
		}
		if art.From != "" {
			return nil, errors.Errorf(errors.CodeBadRequest, "templates.%s.%s.from not valid in inputs", tmpl.Name, artRef)
		}
		errPrefix := fmt.Sprintf("templates.%s.%s", tmpl.Name, artRef)
		err = validateArtifactLocation(errPrefix, art)
		if err != nil {
			return nil, err
		}
	}
	return scope, nil
}

func validateArtifactLocation(errPrefix string, art wfv1.Artifact) error {
	if art.Git != nil {
		if art.Git.Repo == "" {
			return errors.Errorf(errors.CodeBadRequest, "%s.git.repo is required", errPrefix)
		}
	}
	// TODO: validate other artifact locations
	return nil
}

// resolveAllVariables is a helper to ensure all {{variables}} are resolveable from current scope
func resolveAllVariables(scope map[string]interface{}, tmplStr string) error {
	var unresolvedErr error
	_, allowAllItemRefs := scope[anyItemMagicValue] // 'item.*' is a magic placeholder value set by addItemsToScope
	fstTmpl := fasttemplate.New(tmplStr, "{{", "}}")

	fstTmpl.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		_, ok := scope[tag]
		if !ok && unresolvedErr == nil {
			if (tag == "item" || strings.HasPrefix(tag, "item.")) && allowAllItemRefs {
				// we are *probably* referencing a undetermined item using withParam
				// NOTE: this is far from foolproof.
			} else {
				unresolvedErr = fmt.Errorf("failed to resolve {{%s}}", tag)
			}
		}
		return 0, nil
	})
	return unresolvedErr
}

func validateNonLeaf(tmpl *wfv1.Template) error {
	if tmpl.ActiveDeadlineSeconds != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s.activeDeadlineSeconds is only valid for leaf templates", tmpl.Name)
	}
	if tmpl.RetryStrategy != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s.retryStrategy is only valid for container templates", tmpl.Name)
	}
	return nil
}

func validateLeaf(scope map[string]interface{}, tmpl *wfv1.Template) error {
	tmplBytes, err := json.Marshal(tmpl)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	err = resolveAllVariables(scope, string(tmplBytes))
	if err != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s: %s", tmpl.Name, err.Error())
	}
	if tmpl.Container != nil {
		// Ensure there are no collisions with volume mountPaths and artifact load paths
		mountPaths := make(map[string]string)
		for i, volMount := range tmpl.Container.VolumeMounts {
			if prev, ok := mountPaths[volMount.MountPath]; ok {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.container.volumeMounts[%d].mountPath '%s' already mounted in %s", tmpl.Name, i, volMount.MountPath, prev)
			}
			mountPaths[volMount.MountPath] = fmt.Sprintf("container.volumeMounts.%s", volMount.Name)
		}
		for i, art := range tmpl.Inputs.Artifacts {
			if prev, ok := mountPaths[art.Path]; ok {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.inputs.artifacts[%d].path '%s' already mounted in %s", tmpl.Name, i, art.Path, prev)
			}
			mountPaths[art.Path] = fmt.Sprintf("inputs.artifacts.%s", art.Name)
		}
	}
	if tmpl.ActiveDeadlineSeconds != nil {
		if *tmpl.ActiveDeadlineSeconds <= 0 {
			return errors.Errorf(errors.CodeBadRequest, "templates.%s.activeDeadlineSeconds must be a positive integer > 0", tmpl.Name)
		}
	}
	if tmpl.Parallelism != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s.parallelism is only valid for steps and dag templates", tmpl.Name)
	}
	return nil
}

func validateArguments(prefix string, arguments wfv1.Arguments) error {
	fieldToSlices := map[string]interface{}{
		"parameters": arguments.Parameters,
		"artifacts":  arguments.Artifacts,
	}
	for fieldName, lst := range fieldToSlices {
		err := validateWorkflowFieldNames(lst)
		if err != nil {
			return errors.Errorf(errors.CodeBadRequest, "%s%s%s", prefix, fieldName, err.Error())
		}
	}
	for _, param := range arguments.Parameters {
		if param.Value == nil {
			return errors.Errorf(errors.CodeBadRequest, "%svalue is required", prefix)
		}
	}

	for _, art := range arguments.Artifacts {
		if art.From == "" && !art.HasLocation() {
			return errors.Errorf(errors.CodeBadRequest, "%sfrom or artifact location is required", prefix)
		}
	}
	return nil
}

func (ctx *wfValidationCtx) validateSteps(scope map[string]interface{}, tmpl *wfv1.Template) error {
	err := validateNonLeaf(tmpl)
	if err != nil {
		return err
	}
	stepNames := make(map[string]bool)
	for i, stepGroup := range tmpl.Steps {
		for _, step := range stepGroup {
			if step.Name == "" {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.steps[%d].name is required", tmpl.Name, i)
			}
			_, ok := stepNames[step.Name]
			if ok {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.steps[%d].name '%s' is not unique", tmpl.Name, i, step.Name)
			}
			if errs := IsValidWorkflowFieldName(step.Name); len(errs) != 0 {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.steps[%d].name '%s' is invalid: %s", tmpl.Name, i, step.Name, strings.Join(errs, ";"))
			}
			stepNames[step.Name] = true
			err := addItemsToScope(&step, scope)
			if err != nil {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.steps[%d].%s %s", tmpl.Name, i, step.Name, err.Error())
			}
			stepBytes, err := json.Marshal(stepGroup)
			if err != nil {
				return errors.InternalWrapError(err)
			}
			err = resolveAllVariables(scope, string(stepBytes))
			if err != nil {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.steps[%d].%s %s", tmpl.Name, i, step.Name, err.Error())
			}
			childTmpl := ctx.wf.GetTemplate(step.Template)
			if childTmpl == nil {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.steps[%d].%s.template '%s' undefined", tmpl.Name, i, step.Name, step.Template)
			}
			err = validateArguments(fmt.Sprintf("templates.%s.steps[%d].%s.arguments.", tmpl.Name, i, step.Name), step.Arguments)
			if err != nil {
				return err
			}
			err = ctx.validateTemplate(childTmpl, step.Arguments)
			if err != nil {
				return err
			}
		}
		for _, step := range stepGroup {
			ctx.addOutputsToScope(step.Template, fmt.Sprintf("steps.%s", step.Name), scope)
		}
	}
	return nil
}

func addItemsToScope(step *wfv1.WorkflowStep, scope map[string]interface{}) error {
	if len(step.WithItems) > 0 && step.WithParam != "" {
		return fmt.Errorf("only one of withItems or withParam can be specified")
	}
	if len(step.WithItems) > 0 {
		for i := range step.WithItems {
			switch val := step.WithItems[i].Value.(type) {
			case string, int32, int64, float32, float64, bool:
				scope["item"] = true
			case map[string]interface{}:
				for itemKey := range val {
					scope[fmt.Sprintf("item.%s", itemKey)] = true
				}
			default:
				return fmt.Errorf("unsupported withItems type: %v", val)
			}
		}
	} else if step.WithParam != "" {
		scope["item"] = true
		// 'item.*' is magic placeholder value which resolveAllVariables() will look for
		// when considering if all variables are resolveable.
		scope[anyItemMagicValue] = true
	}
	return nil
}

func (ctx *wfValidationCtx) addOutputsToScope(templateName string, prefix string, scope map[string]interface{}) {
	tmpl := ctx.wf.GetTemplate(templateName)
	if tmpl.Daemon != nil && *tmpl.Daemon {
		scope[fmt.Sprintf("%s.ip", prefix)] = true
	}
	if tmpl.Script != nil {
		scope[fmt.Sprintf("%s.outputs.result", prefix)] = true
	}
	for _, param := range tmpl.Outputs.Parameters {
		scope[fmt.Sprintf("%s.outputs.parameters.%s", prefix, param.Name)] = true
	}
	for _, art := range tmpl.Outputs.Artifacts {
		scope[fmt.Sprintf("%s.outputs.artifacts.%s", prefix, art.Name)] = true
	}
}

func validateOutputs(scope map[string]interface{}, tmpl *wfv1.Template) error {
	err := validateWorkflowFieldNames(tmpl.Outputs.Parameters)
	if err != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s.outputs.parameters%s", tmpl.Name, err.Error())
	}
	err = validateWorkflowFieldNames(tmpl.Outputs.Artifacts)
	if err != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s.outputs.artifacts%s", tmpl.Name, err.Error())
	}
	outputBytes, err := json.Marshal(tmpl.Outputs)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	err = resolveAllVariables(scope, string(outputBytes))
	if err != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s.outputs %s", tmpl.Name, err.Error())
	}

	isLeaf := tmpl.Container != nil || tmpl.Script != nil
	for _, art := range tmpl.Outputs.Artifacts {
		artRef := fmt.Sprintf("outputs.artifacts.%s", art.Name)
		if isLeaf {
			if art.Path == "" {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.%s.path not specified", tmpl.Name, artRef)
			}
		} else {
			if art.Path != "" {
				return errors.Errorf(errors.CodeBadRequest, "templates.%s.%s.path only valid in container/script templates", tmpl.Name, artRef)
			}
		}
	}
	for _, param := range tmpl.Outputs.Parameters {
		paramRef := fmt.Sprintf("templates.%s.outputs.parameters.%s", tmpl.Name, param.Name)
		err = validateOutputParameter(paramRef, &param)
		if err != nil {
			return err
		}
		tmplType := tmpl.GetType()
		switch tmplType {
		case wfv1.TemplateTypeContainer, wfv1.TemplateTypeScript:
			if param.ValueFrom.Path == "" {
				return errors.Errorf(errors.CodeBadRequest, "%s.path must be specified for %s templates", paramRef, tmplType)
			}
		case wfv1.TemplateTypeResource:
			if param.ValueFrom.JQFilter == "" && param.ValueFrom.JSONPath == "" {
				return errors.Errorf(errors.CodeBadRequest, "%s .jqFilter or jsonPath must be specified for %s templates", paramRef, tmplType)
			}
		case wfv1.TemplateTypeDAG, wfv1.TemplateTypeSteps:
			if param.ValueFrom.Parameter == "" {
				return errors.Errorf(errors.CodeBadRequest, "%s.parameter must be specified for %s templates", paramRef, tmplType)
			}
		}
	}
	return nil
}

// validateOutputParameter verifies that only one of valueFrom is defined in an output
func validateOutputParameter(paramRef string, param *wfv1.Parameter) error {
	if param.ValueFrom == nil {
		return errors.Errorf(errors.CodeBadRequest, "%s.valueFrom not specified", paramRef)
	}
	paramTypes := 0
	for _, value := range []string{param.ValueFrom.Path, param.ValueFrom.JQFilter, param.ValueFrom.JSONPath, param.ValueFrom.Parameter} {
		if value != "" {
			paramTypes++
		}
	}
	switch paramTypes {
	case 0:
		return errors.New(errors.CodeBadRequest, "valueFrom type unspecified. choose one of: path, jqFilter, jsonPath, parameter")
	case 1:
	default:
		return errors.New(errors.CodeBadRequest, "multiple valueFrom types specified. choose one of: path, jqFilter, jsonPath, parameter")
	}
	return nil
}

// validateWorkflowFieldNames accepts a slice of structs and
// verifies that the Name field of the structs are:
// * unique
// * non-empty
// * matches matches our regex requirements
func validateWorkflowFieldNames(slice interface{}) error {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return errors.InternalErrorf("validateWorkflowFieldNames given a non-slice type")
	}
	items := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		items[i] = s.Index(i).Interface()
	}
	names := make(map[string]bool)
	getNameFieldValue := func(val interface{}) (string, error) {
		s := reflect.ValueOf(val)
		for i := 0; i < s.NumField(); i++ {
			typeField := s.Type().Field(i)
			if typeField.Name == "Name" {
				return s.Field(i).String(), nil
			}
		}
		return "", errors.InternalError("No 'Name' field in struct")
	}

	for i, item := range items {
		name, err := getNameFieldValue(item)
		if err != nil {
			return err
		}
		if name == "" {
			return errors.Errorf(errors.CodeBadRequest, "[%d].name is required", i)
		}
		if errs := IsValidWorkflowFieldName(name); len(errs) != 0 {
			return errors.Errorf(errors.CodeBadRequest, "[%d].name: '%s' is invalid: %s", i, name, strings.Join(errs, ";"))
		}
		_, ok := names[name]
		if ok {
			return errors.Errorf(errors.CodeBadRequest, "[%d].name '%s' is not unique", i, name)
		}
		names[name] = true
	}
	return nil
}

func (ctx *wfValidationCtx) validateDAG(scope map[string]interface{}, tmpl *wfv1.Template) error {
	err := validateNonLeaf(tmpl)
	if err != nil {
		return err
	}
	err = validateWorkflowFieldNames(tmpl.DAG.Tasks)
	if err != nil {
		return errors.Errorf(errors.CodeBadRequest, "templates.%s.tasks%s", tmpl.Name, err.Error())
	}
	nameToTask := make(map[string]wfv1.DAGTask)
	for _, task := range tmpl.DAG.Tasks {
		nameToTask[task.Name] = task
	}

	// Verify dependencies for all tasks can be resolved as well as template names
	for _, task := range tmpl.DAG.Tasks {
		if task.Template == "" {
			return errors.Errorf(errors.CodeBadRequest, "templates.%s.tasks.%s.template is required", tmpl.Name, task.Name)
		}
		taskTmpl := ctx.wf.GetTemplate(task.Template)
		if taskTmpl == nil {
			return errors.Errorf(errors.CodeBadRequest, "templates.%s.tasks%s.template '%s' undefined", tmpl.Name, task.Name, task.Template)
		}
		dupDependencies := make(map[string]bool)
		for j, depName := range task.Dependencies {
			if _, ok := dupDependencies[depName]; ok {
				return errors.Errorf(errors.CodeBadRequest,
					"templates.%s.tasks.%s.dependencies[%d] dependency '%s' duplicated",
					tmpl.Name, task.Name, j, depName)
			}
			dupDependencies[depName] = true
			if _, ok := nameToTask[depName]; !ok {
				return errors.Errorf(errors.CodeBadRequest,
					"templates.%s.tasks.%s.dependencies[%d] dependency '%s' not defined",
					tmpl.Name, task.Name, j, depName)
			}
		}
	}

	err = verifyNoCycles(tmpl, nameToTask)
	if err != nil {
		return err
	}

	for _, task := range tmpl.DAG.Tasks {
		// add all tasks outputs to scope so that DAGs can have outputs
		ctx.addOutputsToScope(task.Template, fmt.Sprintf("tasks.%s", task.Name), scope)

		taskBytes, err := json.Marshal(task)
		if err != nil {
			return errors.InternalWrapError(err)
		}
		taskScope := make(map[string]interface{})
		for k, v := range scope {
			taskScope[k] = v
		}
		ancestry := GetTaskAncestry(task.Name, tmpl.DAG.Tasks)
		for _, ancestor := range ancestry {
			ctx.addOutputsToScope(nameToTask[ancestor].Template, fmt.Sprintf("tasks.%s", ancestor), taskScope)
		}
		err = resolveAllVariables(taskScope, string(taskBytes))
		if err != nil {
			return errors.Errorf(errors.CodeBadRequest, "templates.%s.tasks.%s %s", tmpl.Name, task.Name, err.Error())
		}
		err = validateArguments(fmt.Sprintf("templates.%s.tasks.%s.arguments.", tmpl.Name, task.Name), task.Arguments)
		if err != nil {
			return err
		}
		taskTmpl := ctx.wf.GetTemplate(task.Template)
		err = ctx.validateTemplate(taskTmpl, task.Arguments)
		if err != nil {
			return err
		}
	}

	return nil
}

// verifyNoCycles verifies there are no cycles in the DAG graph
func verifyNoCycles(tmpl *wfv1.Template, nameToTask map[string]wfv1.DAGTask) error {
	visited := make(map[string]bool)
	var noCyclesHelper func(taskName string, cycle []string) error
	noCyclesHelper = func(taskName string, cycle []string) error {
		if _, ok := visited[taskName]; ok {
			return nil
		}
		task := nameToTask[taskName]
		for _, depName := range task.Dependencies {
			for _, name := range cycle {
				if name == depName {
					return errors.Errorf(errors.CodeBadRequest,
						"templates.%s.tasks dependency cycle detected: %s->%s",
						tmpl.Name, strings.Join(cycle, "->"), name)
				}
			}
			cycle = append(cycle, depName)
			err := noCyclesHelper(depName, cycle)
			if err != nil {
				return err
			}
			cycle = cycle[0 : len(cycle)-1]
		}
		visited[taskName] = true
		return nil
	}

	for _, task := range tmpl.DAG.Tasks {
		err := noCyclesHelper(task.Name, []string{})
		if err != nil {
			return err
		}
	}
	return nil
}
