package commands

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"text/tabwriter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	"github.com/argoproj/argo/cmd/argo/commands/common"
	"github.com/argoproj/argo/cmd/argo/commands/test"
	clientmocks "github.com/argoproj/argo/pkg/apiclient/mocks"
	"github.com/argoproj/argo/pkg/apiclient/workflow/mocks"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var wfWithStatus = `
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  creationTimestamp: "2020-06-24T22:53:35Z"
  generation: 6
  labels:
    workflows.argoproj.io/completed: "true"
    workflows.argoproj.io/phase: Succeeded
  name: hello-world
  namespace: default
  resourceVersion: "1110858"
  selfLink: /apis/argoproj.io/v1alpha1/namespaces/default/workflows/hello-world-2xg9p
  uid: 8c25e2e7-6b35-4a49-a667-87b4cd1afa3c
spec:
  arguments: {}
  entrypoint: whalesay
  templates:
  - arguments: {}
    container:
      args:
      - hello world
      command:
      - cowsay
      image: docker/whalesay:latest
      name: ""
      resources: {}
    inputs: {}
    metadata: {}
    name: whalesay
    outputs: {}
status:
  conditions:
  - status: "True"
    type: Completed
  finishedAt: "2020-06-24T22:53:41Z"
  nodes:
    hello-world-2xg9p:
      displayName: hello-world-2xg9p
      finishedAt: "2020-06-24T22:53:39Z"
      id: hello-world-2xg9p
      name: hello-world-2xg9p
      outputs:
        exitCode: "0"
      phase: Succeeded
      resourcesDuration:
        cpu: 3
        memory: 0
      startedAt: "2020-06-24T22:53:35Z"
      templateName: whalesay
      templateScope: local/hello-world-2xg9p
      type: Pod
  phase: Succeeded
  resourcesDuration:
    cpu: 3
    memory: 0
  startedAt: "2020-06-24T22:53:35Z"
`

func testPrintNodeImpl(t *testing.T, expected string, node wfv1.NodeStatus, nodePrefix string, getArgs getFlags) {
	var result bytes.Buffer
	w := tabwriter.NewWriter(&result, 0, 8, 1, '\t', 0)
	filtered, _ := filterNode(node, getArgs)
	if !filtered {
		printNode(w, node, nodePrefix, getArgs)
	}
	err := w.Flush()
	assert.NoError(t, err)
	assert.Equal(t, expected, result.String())
}

// TestPrintNode
func TestPrintNode(t *testing.T) {
	nodeName := "testNode"
	kubernetesNodeName := "testKnodeName"
	nodePrefix := ""
	nodeTemplateName := "testTemplate"
	nodeTemplateRefName := "testTemplateRef"
	nodeID := "testID"
	nodeMessage := "test"
	getArgs := getFlags{
		output: "",
	}
	timestamp := metav1.Time{
		Time: time.Now(),
	}
	node := wfv1.NodeStatus{
		Name:        nodeName,
		Phase:       wfv1.NodeRunning,
		DisplayName: nodeName,
		Type:        wfv1.NodeTypePod,
		ID:          nodeID,
		StartedAt:   timestamp,
		FinishedAt:  timestamp,
		Message:     nodeMessage,
	}
	node.HostNodeName = kubernetesNodeName
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s\t%s\t%s\t%s\t%s\n", common.JobStatusIconMap[wfv1.NodeRunning], nodeName, "", nodeID, "0s", nodeMessage, ""), node, nodePrefix, getArgs)

	// Compatibility test
	getArgs.status = "Running"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t\t%s\t%s\t%s\t\n", common.JobStatusIconMap[wfv1.NodeRunning], nodeName, nodeID, "0s", nodeMessage), node, nodePrefix, getArgs)

	getArgs.status = ""
	getArgs.nodeFieldSelectorString = "phase=Running"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t\t%s\t%s\t%s\t\n", common.JobStatusIconMap[wfv1.NodeRunning], nodeName, nodeID, "0s", nodeMessage), node, nodePrefix, getArgs)

	getArgs.nodeFieldSelectorString = "phase!=foobar"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t\t%s\t%s\t%s\t\n", common.JobStatusIconMap[wfv1.NodeRunning], nodeName, nodeID, "0s", nodeMessage), node, nodePrefix, getArgs)

	getArgs.nodeFieldSelectorString = "phase!=Running"
	testPrintNodeImpl(t, "", node, nodePrefix, getArgs)

	// Compatibility test
	getArgs.nodeFieldSelectorString = ""
	getArgs.status = "foobar"
	testPrintNodeImpl(t, "", node, nodePrefix, getArgs)

	getArgs.status = ""
	getArgs.nodeFieldSelectorString = "phase=foobar"
	testPrintNodeImpl(t, "", node, nodePrefix, getArgs)

	getArgs = getFlags{
		output: "",
	}

	node.TemplateName = nodeTemplateName
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s\t%s\t%s\t%s\t%s\n", common.JobStatusIconMap[wfv1.NodeRunning], nodeName, nodeTemplateName, nodeID, "0s", nodeMessage, ""), node, nodePrefix, getArgs)

	node.Type = wfv1.NodeTypeSuspend
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s\t%s\t%s\t%s\t%s\n", common.NodeTypeIconMap[wfv1.NodeTypeSuspend], nodeName, nodeTemplateName, "", "", nodeMessage, ""), node, nodePrefix, getArgs)

	node.TemplateRef = &wfv1.TemplateRef{
		Name:     nodeTemplateRefName,
		Template: nodeTemplateRefName,
	}
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s/%s\t%s\t%s\t%s\t%s\n", common.NodeTypeIconMap[wfv1.NodeTypeSuspend], nodeName, nodeTemplateRefName, nodeTemplateRefName, "", "", nodeMessage, ""), node, nodePrefix, getArgs)

	getArgs.output = "wide"
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s/%s\t%s\t%s\t%s\t%s\t%s\t\n", common.NodeTypeIconMap[wfv1.NodeTypeSuspend], nodeName, nodeTemplateRefName, nodeTemplateRefName, "", "", getArtifactsString(node), nodeMessage, ""), node, nodePrefix, getArgs)

	node.Type = wfv1.NodeTypePod
	testPrintNodeImpl(t, fmt.Sprintf("%s %s\t%s/%s\t%s\t%s\t%s\t%s\t%s\t%s\n", common.JobStatusIconMap[wfv1.NodeRunning], nodeName, nodeTemplateRefName, nodeTemplateRefName, nodeID, "0s", getArtifactsString(node), nodeMessage, "", kubernetesNodeName), node, nodePrefix, getArgs)

	getArgs.status = "foobar"
	testPrintNodeImpl(t, "", node, nodePrefix, getArgs)
}

func TestStatusToNodeFieldSelector(t *testing.T) {
	one := statusToNodeFieldSelector("Running")
	assert.Equal(t, "phase=Running", one)
}

var indexTest = `apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  creationTimestamp: "2020-06-02T16:04:21Z"
  generateName: many-items-
  generation: 32
  labels:
    workflows.argoproj.io/completed: "true"
    workflows.argoproj.io/phase: Succeeded
  name: many-items-z26lj
  namespace: argo
  resourceVersion: "5102"
  selfLink: /apis/argoproj.io/v1alpha1/namespaces/argo/workflows/many-items-z26lj
  uid: d21f092a-f659-4300-bd69-983a9912a379
spec:
  arguments: {}
  entrypoint: parallel-sleep
  templates:
  - arguments: {}
    inputs: {}
    metadata: {}
    name: parallel-sleep
    outputs: {}
    steps:
    - - arguments: {}
        name: sleep
        template: sleep
        withItems:
        - zero
        - one
        - two
        - three
        - four
        - five
        - six
        - seven
        - eight
        - nine
        - ten
        - eleven
        - twelve
  - arguments: {}
    container:
      command:
      - sh
      - -c
      - sleep 10
      image: alpine:latest
      name: ""
      resources: {}
    inputs: {}
    metadata: {}
    name: sleep
    outputs: {}
status:
  conditions:
  - status: "True"
    type: Completed
  finishedAt: "2020-06-02T16:05:01Z"
  nodes:
    many-items-z26lj:
      children:
      - many-items-z26lj-1414877240
      displayName: many-items-z26lj
      finishedAt: "2020-06-02T16:05:01Z"
      id: many-items-z26lj
      name: many-items-z26lj
      outboundNodes:
      - many-items-z26lj-1939921510
      - many-items-z26lj-2156977535
      - many-items-z26lj-3409403178
      - many-items-z26lj-1774150289
      - many-items-z26lj-3491220632
      - many-items-z26lj-1942531647
      - many-items-z26lj-3178865096
      - many-items-z26lj-3031375822
      - many-items-z26lj-753834747
      - many-items-z26lj-2619926859
      - many-items-z26lj-1052882686
      - many-items-z26lj-3011405271
      - many-items-z26lj-3126938806
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: parallel-sleep
      type: Steps
    many-items-z26lj-753834747:
      boundaryID: many-items-z26lj
      displayName: sleep(8:eight)
      finishedAt: "2020-06-02T16:04:42Z"
      id: many-items-z26lj-753834747
      name: many-items-z26lj[0].sleep(8:eight)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1052882686:
      boundaryID: many-items-z26lj
      displayName: sleep(10:ten)
      finishedAt: "2020-06-02T16:04:45Z"
      id: many-items-z26lj-1052882686
      name: many-items-z26lj[0].sleep(10:ten)
      startedAt: "2020-06-02T16:04:22Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1414877240:
      boundaryID: many-items-z26lj
      children:
      - many-items-z26lj-1939921510
      - many-items-z26lj-2156977535
      - many-items-z26lj-3409403178
      - many-items-z26lj-1774150289
      - many-items-z26lj-3491220632
      - many-items-z26lj-1942531647
      - many-items-z26lj-3178865096
      - many-items-z26lj-3031375822
      - many-items-z26lj-753834747
      - many-items-z26lj-2619926859
      - many-items-z26lj-1052882686
      - many-items-z26lj-3011405271
      - many-items-z26lj-3126938806
      displayName: '[0]'
      finishedAt: "2020-06-02T16:05:01Z"
      id: many-items-z26lj-1414877240
      name: many-items-z26lj[0]
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: parallel-sleep
      type: StepGroup
    many-items-z26lj-1774150289:
      boundaryID: many-items-z26lj
      displayName: sleep(3:three)
      finishedAt: "2020-06-02T16:04:54Z"
      id: many-items-z26lj-1774150289
      name: many-items-z26lj[0].sleep(3:three)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1939921510:
      boundaryID: many-items-z26lj
      displayName: sleep(0:zero)
      finishedAt: "2020-06-02T16:04:48Z"
      id: many-items-z26lj-1939921510
      name: many-items-z26lj[0].sleep(0:zero)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-1942531647:
      boundaryID: many-items-z26lj
      displayName: sleep(5:five)
      finishedAt: "2020-06-02T16:04:47Z"
      id: many-items-z26lj-1942531647
      name: many-items-z26lj[0].sleep(5:five)
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-2156977535:
      boundaryID: many-items-z26lj
      displayName: sleep(1:one)
      finishedAt: "2020-06-02T16:04:53Z"
      id: many-items-z26lj-2156977535
      name: many-items-z26lj[0].sleep(1:one)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-2619926859:
      boundaryID: many-items-z26lj
      displayName: sleep(9:nine)
      finishedAt: "2020-06-02T16:04:40Z"
      id: many-items-z26lj-2619926859
      name: many-items-z26lj[0].sleep(9:nine)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3011405271:
      boundaryID: many-items-z26lj
      displayName: sleep(11:eleven)
      finishedAt: "2020-06-02T16:04:44Z"
      id: many-items-z26lj-3011405271
      name: many-items-z26lj[0].sleep(11:eleven)
      startedAt: "2020-06-02T16:04:22Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3031375822:
      boundaryID: many-items-z26lj
      displayName: sleep(7:seven)
      finishedAt: "2020-06-02T16:04:57Z"
      id: many-items-z26lj-3031375822
      name: many-items-z26lj[0].sleep(7:seven)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3126938806:
      boundaryID: many-items-z26lj
      displayName: sleep(12:twelve)
      finishedAt: "2020-06-02T16:04:59Z"
      id: many-items-z26lj-3126938806
      name: many-items-z26lj[0].sleep(12:twelve)
      startedAt: "2020-06-02T16:04:22Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3178865096:
      boundaryID: many-items-z26lj
      displayName: sleep(6:six)
      finishedAt: "2020-06-02T16:04:56Z"
      id: many-items-z26lj-3178865096
      name: many-items-z26lj[0].sleep(6:six)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3409403178:
      boundaryID: many-items-z26lj
      displayName: sleep(2:two)
      finishedAt: "2020-06-02T16:04:51Z"
      id: many-items-z26lj-3409403178
      name: many-items-z26lj[0].sleep(2:two)
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
    many-items-z26lj-3491220632:
      boundaryID: many-items-z26lj
      displayName: sleep(4:four)
      finishedAt: "2020-06-02T16:04:50Z"
      id: many-items-z26lj-3491220632
      name: many-items-z26lj[0].sleep(4:four)
      phase: Succeeded
      startedAt: "2020-06-02T16:04:21Z"
      templateName: sleep
      type: Pod
  phase: Succeeded
  startedAt: "2020-06-02T16:04:21Z"
`

func TestIndexOrdering(t *testing.T) {
	var wf wfv1.Workflow
	err := yaml.Unmarshal([]byte(indexTest), &wf)
	if assert.NoError(t, err) {
		assert.Contains(t, printWorkflowHelper(&wf, getFlags{}), `         
   ├- sleep(9:nine)     sleep           many-items-z26lj-2619926859  19s         
   ├- sleep(10:ten)     sleep           many-items-z26lj-1052882686  23s         
   ├- sleep(11:eleven)  sleep           many-items-z26lj-3011405271  22s`)
	}
}

func TestGetCommand(t *testing.T) {
	getOutput := `Name:                hello-world
Namespace:           default
ServiceAccount:      default
Status:              Succeeded
`
	var wf wfv1.Workflow
	err := yaml.Unmarshal([]byte(wfWithStatus), &wf)
	assert.NoError(t, err)
	client := clientmocks.Client{}
	wfClient := mocks.WorkflowServiceClient{}
	wfClient.On("GetWorkflow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&wf, nil)
	client.On("NewWorkflowServiceClient").Return(&wfClient)
	common.APIClient = &client
	getCommand := NewGetCommand()
	getCommand.SetArgs([]string{"hello-world"})
	execFunc := func() {
		err := getCommand.Execute()
		assert.NoError(t, err)
	}
	output := test.CaptureOutput(execFunc)
	assert.Contains(t, output, getOutput)
}
