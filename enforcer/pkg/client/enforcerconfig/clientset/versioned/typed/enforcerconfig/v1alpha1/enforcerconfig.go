/*
Copyright The 2020 IBM Corporation.

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

	v1alpha1 "github.com/IBM/integrity-enforcer/enforcer/pkg/apis/enforcerconfig/v1alpha1"
	scheme "github.com/IBM/integrity-enforcer/enforcer/pkg/client/enforcerconfig/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnforcerConfigsGetter has a method to return a EnforcerConfigInterface.
// A group's client should implement this interface.
type EnforcerConfigsGetter interface {
	EnforcerConfigs(namespace string) EnforcerConfigInterface
}

// EnforcerConfigInterface has methods to work with EnforcerConfig resources.
type EnforcerConfigInterface interface {
	Create(*v1alpha1.EnforcerConfig) (*v1alpha1.EnforcerConfig, error)
	Update(*v1alpha1.EnforcerConfig) (*v1alpha1.EnforcerConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EnforcerConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.EnforcerConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EnforcerConfig, err error)
	EnforcerConfigExpansion
}

// enforcerConfigs implements EnforcerConfigInterface
type enforcerConfigs struct {
	client rest.Interface
	ns     string
}

// newEnforcerConfigs returns a EnforcerConfigs
func newEnforcerConfigs(c *ResearchV1alpha1Client, namespace string) *enforcerConfigs {
	return &enforcerConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the enforcerConfig, and returns the corresponding enforcerConfig object, and an error if there is any.
func (c *enforcerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.EnforcerConfig, err error) {
	result = &v1alpha1.EnforcerConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("enforcerconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EnforcerConfigs that match those selectors.
func (c *enforcerConfigs) List(opts v1.ListOptions) (result *v1alpha1.EnforcerConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EnforcerConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("enforcerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested enforcerConfigs.
func (c *enforcerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("enforcerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a enforcerConfig and creates it.  Returns the server's representation of the enforcerConfig, and an error, if there is any.
func (c *enforcerConfigs) Create(enforcerConfig *v1alpha1.EnforcerConfig) (result *v1alpha1.EnforcerConfig, err error) {
	result = &v1alpha1.EnforcerConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("enforcerconfigs").
		Body(enforcerConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a enforcerConfig and updates it. Returns the server's representation of the enforcerConfig, and an error, if there is any.
func (c *enforcerConfigs) Update(enforcerConfig *v1alpha1.EnforcerConfig) (result *v1alpha1.EnforcerConfig, err error) {
	result = &v1alpha1.EnforcerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("enforcerconfigs").
		Name(enforcerConfig.Name).
		Body(enforcerConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the enforcerConfig and deletes it. Returns an error if one occurs.
func (c *enforcerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("enforcerconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *enforcerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("enforcerconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched enforcerConfig.
func (c *enforcerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EnforcerConfig, err error) {
	result = &v1alpha1.EnforcerConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("enforcerconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
