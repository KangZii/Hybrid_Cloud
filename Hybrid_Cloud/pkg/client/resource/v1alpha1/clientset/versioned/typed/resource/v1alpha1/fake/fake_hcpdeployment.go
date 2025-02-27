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

package fake

import (
	v1alpha1 "Hybrid_Cloud/pkg/apis/resource/v1alpha1"
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHCPDeployments implements HCPDeploymentInterface
type FakeHCPDeployments struct {
	Fake *FakeHcpV1alpha1
	ns   string
}

var hcpdeploymentsResource = schema.GroupVersionResource{Group: "hcp.crd.com", Version: "v1alpha1", Resource: "hcpdeployments"}

var hcpdeploymentsKind = schema.GroupVersionKind{Group: "hcp.crd.com", Version: "v1alpha1", Kind: "HCPDeployment"}

// Get takes name of the hCPDeployment, and returns the corresponding hCPDeployment object, and an error if there is any.
func (c *FakeHCPDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HCPDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hcpdeploymentsResource, c.ns, name), &v1alpha1.HCPDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HCPDeployment), err
}

// List takes label and field selectors, and returns the list of HCPDeployments that match those selectors.
func (c *FakeHCPDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HCPDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hcpdeploymentsResource, hcpdeploymentsKind, c.ns, opts), &v1alpha1.HCPDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HCPDeploymentList{ListMeta: obj.(*v1alpha1.HCPDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.HCPDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hCPDeployments.
func (c *FakeHCPDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hcpdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a hCPDeployment and creates it.  Returns the server's representation of the hCPDeployment, and an error, if there is any.
func (c *FakeHCPDeployments) Create(ctx context.Context, hCPDeployment *v1alpha1.HCPDeployment, opts v1.CreateOptions) (result *v1alpha1.HCPDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hcpdeploymentsResource, c.ns, hCPDeployment), &v1alpha1.HCPDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HCPDeployment), err
}

// Update takes the representation of a hCPDeployment and updates it. Returns the server's representation of the hCPDeployment, and an error, if there is any.
func (c *FakeHCPDeployments) Update(ctx context.Context, hCPDeployment *v1alpha1.HCPDeployment, opts v1.UpdateOptions) (result *v1alpha1.HCPDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hcpdeploymentsResource, c.ns, hCPDeployment), &v1alpha1.HCPDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HCPDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHCPDeployments) UpdateStatus(ctx context.Context, hCPDeployment *v1alpha1.HCPDeployment, opts v1.UpdateOptions) (*v1alpha1.HCPDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hcpdeploymentsResource, "status", c.ns, hCPDeployment), &v1alpha1.HCPDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HCPDeployment), err
}

// Delete takes name of the hCPDeployment and deletes it. Returns an error if one occurs.
func (c *FakeHCPDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hcpdeploymentsResource, c.ns, name, opts), &v1alpha1.HCPDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHCPDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hcpdeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HCPDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched hCPDeployment.
func (c *FakeHCPDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HCPDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hcpdeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HCPDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HCPDeployment), err
}
