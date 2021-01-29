// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/argoproj/argo/v3/pkg/apis/workflow/v1alpha1"
	scheme "github.com/argoproj/argo/v3/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CronWorkflowsGetter has a method to return a CronWorkflowInterface.
// A group's client should implement this interface.
type CronWorkflowsGetter interface {
	CronWorkflows(namespace string) CronWorkflowInterface
}

// CronWorkflowInterface has methods to work with CronWorkflow resources.
type CronWorkflowInterface interface {
	Create(ctx context.Context, cronWorkflow *v1alpha1.CronWorkflow, opts v1.CreateOptions) (*v1alpha1.CronWorkflow, error)
	Update(ctx context.Context, cronWorkflow *v1alpha1.CronWorkflow, opts v1.UpdateOptions) (*v1alpha1.CronWorkflow, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CronWorkflow, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CronWorkflowList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CronWorkflow, err error)
	CronWorkflowExpansion
}

// cronWorkflows implements CronWorkflowInterface
type cronWorkflows struct {
	client rest.Interface
	ns     string
}

// newCronWorkflows returns a CronWorkflows
func newCronWorkflows(c *ArgoprojV1alpha1Client, namespace string) *cronWorkflows {
	return &cronWorkflows{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cronWorkflow, and returns the corresponding cronWorkflow object, and an error if there is any.
func (c *cronWorkflows) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CronWorkflow, err error) {
	result = &v1alpha1.CronWorkflow{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cronworkflows").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CronWorkflows that match those selectors.
func (c *cronWorkflows) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CronWorkflowList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CronWorkflowList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cronworkflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cronWorkflows.
func (c *cronWorkflows) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cronworkflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cronWorkflow and creates it.  Returns the server's representation of the cronWorkflow, and an error, if there is any.
func (c *cronWorkflows) Create(ctx context.Context, cronWorkflow *v1alpha1.CronWorkflow, opts v1.CreateOptions) (result *v1alpha1.CronWorkflow, err error) {
	result = &v1alpha1.CronWorkflow{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cronworkflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cronWorkflow).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cronWorkflow and updates it. Returns the server's representation of the cronWorkflow, and an error, if there is any.
func (c *cronWorkflows) Update(ctx context.Context, cronWorkflow *v1alpha1.CronWorkflow, opts v1.UpdateOptions) (result *v1alpha1.CronWorkflow, err error) {
	result = &v1alpha1.CronWorkflow{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cronworkflows").
		Name(cronWorkflow.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cronWorkflow).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cronWorkflow and deletes it. Returns an error if one occurs.
func (c *cronWorkflows) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cronworkflows").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cronWorkflows) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cronworkflows").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cronWorkflow.
func (c *cronWorkflows) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CronWorkflow, err error) {
	result = &v1alpha1.CronWorkflow{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cronworkflows").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
