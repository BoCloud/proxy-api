/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/bocloud/proxy-api/apis/proxy/v1beta1"
	scheme "github.com/bocloud/proxy-api/generated/proxy/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TCPRoutesGetter has a method to return a TCPRouteInterface.
// A group's client should implement this interface.
type TCPRoutesGetter interface {
	TCPRoutes(namespace string) TCPRouteInterface
}

// TCPRouteInterface has methods to work with TCPRoute resources.
type TCPRouteInterface interface {
	Create(ctx context.Context, tCPRoute *v1beta1.TCPRoute, opts v1.CreateOptions) (*v1beta1.TCPRoute, error)
	Update(ctx context.Context, tCPRoute *v1beta1.TCPRoute, opts v1.UpdateOptions) (*v1beta1.TCPRoute, error)
	UpdateStatus(ctx context.Context, tCPRoute *v1beta1.TCPRoute, opts v1.UpdateOptions) (*v1beta1.TCPRoute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.TCPRoute, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.TCPRouteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TCPRoute, err error)
	TCPRouteExpansion
}

// tCPRoutes implements TCPRouteInterface
type tCPRoutes struct {
	client rest.Interface
	ns     string
}

// newTCPRoutes returns a TCPRoutes
func newTCPRoutes(c *ProxyV1beta1Client, namespace string) *tCPRoutes {
	return &tCPRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tCPRoute, and returns the corresponding tCPRoute object, and an error if there is any.
func (c *tCPRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.TCPRoute, err error) {
	result = &v1beta1.TCPRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tcproutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TCPRoutes that match those selectors.
func (c *tCPRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.TCPRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.TCPRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tcproutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tCPRoutes.
func (c *tCPRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tcproutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tCPRoute and creates it.  Returns the server's representation of the tCPRoute, and an error, if there is any.
func (c *tCPRoutes) Create(ctx context.Context, tCPRoute *v1beta1.TCPRoute, opts v1.CreateOptions) (result *v1beta1.TCPRoute, err error) {
	result = &v1beta1.TCPRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tcproutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tCPRoute).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tCPRoute and updates it. Returns the server's representation of the tCPRoute, and an error, if there is any.
func (c *tCPRoutes) Update(ctx context.Context, tCPRoute *v1beta1.TCPRoute, opts v1.UpdateOptions) (result *v1beta1.TCPRoute, err error) {
	result = &v1beta1.TCPRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tcproutes").
		Name(tCPRoute.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tCPRoute).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tCPRoutes) UpdateStatus(ctx context.Context, tCPRoute *v1beta1.TCPRoute, opts v1.UpdateOptions) (result *v1beta1.TCPRoute, err error) {
	result = &v1beta1.TCPRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tcproutes").
		Name(tCPRoute.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tCPRoute).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tCPRoute and deletes it. Returns an error if one occurs.
func (c *tCPRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tcproutes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tCPRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tcproutes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tCPRoute.
func (c *tCPRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TCPRoute, err error) {
	result = &v1beta1.TCPRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tcproutes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
