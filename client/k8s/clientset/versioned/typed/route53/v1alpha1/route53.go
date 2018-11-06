/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/tuannvm/r53-operator/apis/route53/v1alpha1"
	scheme "github.com/tuannvm/r53-operator/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// Route53sGetter has a method to return a Route53Interface.
// A group's client should implement this interface.
type Route53sGetter interface {
	Route53s() Route53Interface
}

// Route53Interface has methods to work with Route53 resources.
type Route53Interface interface {
	Create(*v1alpha1.Route53) (*v1alpha1.Route53, error)
	Update(*v1alpha1.Route53) (*v1alpha1.Route53, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Route53, error)
	List(opts v1.ListOptions) (*v1alpha1.Route53List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53, err error)
	Route53Expansion
}

// route53s implements Route53Interface
type route53s struct {
	client rest.Interface
}

// newRoute53s returns a Route53s
func newRoute53s(c *Route53V1alpha1Client) *route53s {
	return &route53s{
		client: c.RESTClient(),
	}
}

// Get takes name of the route53, and returns the corresponding route53 object, and an error if there is any.
func (c *route53s) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53, err error) {
	result = &v1alpha1.Route53{}
	err = c.client.Get().
		Resource("route53s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Route53s that match those selectors.
func (c *route53s) List(opts v1.ListOptions) (result *v1alpha1.Route53List, err error) {
	result = &v1alpha1.Route53List{}
	err = c.client.Get().
		Resource("route53s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested route53s.
func (c *route53s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("route53s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a route53 and creates it.  Returns the server's representation of the route53, and an error, if there is any.
func (c *route53s) Create(route53 *v1alpha1.Route53) (result *v1alpha1.Route53, err error) {
	result = &v1alpha1.Route53{}
	err = c.client.Post().
		Resource("route53s").
		Body(route53).
		Do().
		Into(result)
	return
}

// Update takes the representation of a route53 and updates it. Returns the server's representation of the route53, and an error, if there is any.
func (c *route53s) Update(route53 *v1alpha1.Route53) (result *v1alpha1.Route53, err error) {
	result = &v1alpha1.Route53{}
	err = c.client.Put().
		Resource("route53s").
		Name(route53.Name).
		Body(route53).
		Do().
		Into(result)
	return
}

// Delete takes name of the route53 and deletes it. Returns an error if one occurs.
func (c *route53s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("route53s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *route53s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("route53s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched route53.
func (c *route53s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53, err error) {
	result = &v1alpha1.Route53{}
	err = c.client.Patch(pt).
		Resource("route53s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}