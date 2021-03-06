/*
Copyright 2020 IBM Corporation

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
	"time"

	v1alpha1 "github.com/IBM/integrity-enforcer/enforcer/pkg/apis/signpolicy/v1alpha1"
	scheme "github.com/IBM/integrity-enforcer/enforcer/pkg/client/signpolicy/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SignPoliciesGetter has a method to return a SignPolicyInterface.
// A group's client should implement this interface.
type SignPoliciesGetter interface {
	SignPolicies(namespace string) SignPolicyInterface
}

// SignPolicyInterface has methods to work with SignPolicy resources.
type SignPolicyInterface interface {
	Create(*v1alpha1.SignPolicy) (*v1alpha1.SignPolicy, error)
	Update(*v1alpha1.SignPolicy) (*v1alpha1.SignPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SignPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.SignPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SignPolicy, err error)
	SignPolicyExpansion
}

// signPolicies implements SignPolicyInterface
type signPolicies struct {
	client rest.Interface
	ns     string
}

// newSignPolicies returns a SignPolicies
func newSignPolicies(c *ResearchV1alpha1Client, namespace string) *signPolicies {
	return &signPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the signPolicy, and returns the corresponding signPolicy object, and an error if there is any.
func (c *signPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.SignPolicy, err error) {
	result = &v1alpha1.SignPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("signpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SignPolicies that match those selectors.
func (c *signPolicies) List(opts v1.ListOptions) (result *v1alpha1.SignPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SignPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("signpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested signPolicies.
func (c *signPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("signpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a signPolicy and creates it.  Returns the server's representation of the signPolicy, and an error, if there is any.
func (c *signPolicies) Create(signPolicy *v1alpha1.SignPolicy) (result *v1alpha1.SignPolicy, err error) {
	result = &v1alpha1.SignPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("signpolicies").
		Body(signPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a signPolicy and updates it. Returns the server's representation of the signPolicy, and an error, if there is any.
func (c *signPolicies) Update(signPolicy *v1alpha1.SignPolicy) (result *v1alpha1.SignPolicy, err error) {
	result = &v1alpha1.SignPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("signpolicies").
		Name(signPolicy.Name).
		Body(signPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the signPolicy and deletes it. Returns an error if one occurs.
func (c *signPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("signpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *signPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("signpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched signPolicy.
func (c *signPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SignPolicy, err error) {
	result = &v1alpha1.SignPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("signpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
