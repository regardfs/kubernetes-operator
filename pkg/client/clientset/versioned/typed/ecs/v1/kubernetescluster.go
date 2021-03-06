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

package v1

import (
	"time"

	v1 "github.com/gosoon/kubernetes-operator/pkg/apis/ecs/v1"
	scheme "github.com/gosoon/kubernetes-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KubernetesClustersGetter has a method to return a KubernetesClusterInterface.
// A group's client should implement this interface.
type KubernetesClustersGetter interface {
	KubernetesClusters(namespace string) KubernetesClusterInterface
}

// KubernetesClusterInterface has methods to work with KubernetesCluster resources.
type KubernetesClusterInterface interface {
	Create(*v1.KubernetesCluster) (*v1.KubernetesCluster, error)
	Update(*v1.KubernetesCluster) (*v1.KubernetesCluster, error)
	UpdateStatus(*v1.KubernetesCluster) (*v1.KubernetesCluster, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KubernetesCluster, error)
	List(opts metav1.ListOptions) (*v1.KubernetesClusterList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubernetesCluster, err error)
	KubernetesClusterExpansion
}

// kubernetesClusters implements KubernetesClusterInterface
type kubernetesClusters struct {
	client rest.Interface
	ns     string
}

// newKubernetesClusters returns a KubernetesClusters
func newKubernetesClusters(c *EcsV1Client, namespace string) *kubernetesClusters {
	return &kubernetesClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kubernetesCluster, and returns the corresponding kubernetesCluster object, and an error if there is any.
func (c *kubernetesClusters) Get(name string, options metav1.GetOptions) (result *v1.KubernetesCluster, err error) {
	result = &v1.KubernetesCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KubernetesClusters that match those selectors.
func (c *kubernetesClusters) List(opts metav1.ListOptions) (result *v1.KubernetesClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KubernetesClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubernetesClusters.
func (c *kubernetesClusters) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kubernetesCluster and creates it.  Returns the server's representation of the kubernetesCluster, and an error, if there is any.
func (c *kubernetesClusters) Create(kubernetesCluster *v1.KubernetesCluster) (result *v1.KubernetesCluster, err error) {
	result = &v1.KubernetesCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		Body(kubernetesCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kubernetesCluster and updates it. Returns the server's representation of the kubernetesCluster, and an error, if there is any.
func (c *kubernetesClusters) Update(kubernetesCluster *v1.KubernetesCluster) (result *v1.KubernetesCluster, err error) {
	result = &v1.KubernetesCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		Name(kubernetesCluster.Name).
		Body(kubernetesCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kubernetesClusters) UpdateStatus(kubernetesCluster *v1.KubernetesCluster) (result *v1.KubernetesCluster, err error) {
	result = &v1.KubernetesCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		Name(kubernetesCluster.Name).
		SubResource("status").
		Body(kubernetesCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the kubernetesCluster and deletes it. Returns an error if one occurs.
func (c *kubernetesClusters) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubernetesClusters) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubernetesclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kubernetesCluster.
func (c *kubernetesClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubernetesCluster, err error) {
	result = &v1.KubernetesCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kubernetesclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
