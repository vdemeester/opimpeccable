/*
Copyright 2022 The OpenShift Pipelines Authors

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
	"context"

	v1alpha1 "github.com/vdemeester/opimpeccable/pkg/apis/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenShiftPipelinesConfigs implements OpenShiftPipelinesConfigInterface
type FakeOpenShiftPipelinesConfigs struct {
	Fake *FakeSamplesV1alpha1
	ns   string
}

var openshiftpipelinesconfigsResource = schema.GroupVersionResource{Group: "samples.knative.dev", Version: "v1alpha1", Resource: "openshiftpipelinesconfigs"}

var openshiftpipelinesconfigsKind = schema.GroupVersionKind{Group: "samples.knative.dev", Version: "v1alpha1", Kind: "OpenShiftPipelinesConfig"}

// Get takes name of the openShiftPipelinesConfig, and returns the corresponding openShiftPipelinesConfig object, and an error if there is any.
func (c *FakeOpenShiftPipelinesConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpenShiftPipelinesConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(openshiftpipelinesconfigsResource, c.ns, name), &v1alpha1.OpenShiftPipelinesConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftPipelinesConfig), err
}

// List takes label and field selectors, and returns the list of OpenShiftPipelinesConfigs that match those selectors.
func (c *FakeOpenShiftPipelinesConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpenShiftPipelinesConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(openshiftpipelinesconfigsResource, openshiftpipelinesconfigsKind, c.ns, opts), &v1alpha1.OpenShiftPipelinesConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpenShiftPipelinesConfigList{ListMeta: obj.(*v1alpha1.OpenShiftPipelinesConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpenShiftPipelinesConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openShiftPipelinesConfigs.
func (c *FakeOpenShiftPipelinesConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(openshiftpipelinesconfigsResource, c.ns, opts))

}

// Create takes the representation of a openShiftPipelinesConfig and creates it.  Returns the server's representation of the openShiftPipelinesConfig, and an error, if there is any.
func (c *FakeOpenShiftPipelinesConfigs) Create(ctx context.Context, openShiftPipelinesConfig *v1alpha1.OpenShiftPipelinesConfig, opts v1.CreateOptions) (result *v1alpha1.OpenShiftPipelinesConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(openshiftpipelinesconfigsResource, c.ns, openShiftPipelinesConfig), &v1alpha1.OpenShiftPipelinesConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftPipelinesConfig), err
}

// Update takes the representation of a openShiftPipelinesConfig and updates it. Returns the server's representation of the openShiftPipelinesConfig, and an error, if there is any.
func (c *FakeOpenShiftPipelinesConfigs) Update(ctx context.Context, openShiftPipelinesConfig *v1alpha1.OpenShiftPipelinesConfig, opts v1.UpdateOptions) (result *v1alpha1.OpenShiftPipelinesConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(openshiftpipelinesconfigsResource, c.ns, openShiftPipelinesConfig), &v1alpha1.OpenShiftPipelinesConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftPipelinesConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenShiftPipelinesConfigs) UpdateStatus(ctx context.Context, openShiftPipelinesConfig *v1alpha1.OpenShiftPipelinesConfig, opts v1.UpdateOptions) (*v1alpha1.OpenShiftPipelinesConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(openshiftpipelinesconfigsResource, "status", c.ns, openShiftPipelinesConfig), &v1alpha1.OpenShiftPipelinesConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftPipelinesConfig), err
}

// Delete takes name of the openShiftPipelinesConfig and deletes it. Returns an error if one occurs.
func (c *FakeOpenShiftPipelinesConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(openshiftpipelinesconfigsResource, c.ns, name, opts), &v1alpha1.OpenShiftPipelinesConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenShiftPipelinesConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(openshiftpipelinesconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpenShiftPipelinesConfigList{})
	return err
}

// Patch applies the patch and returns the patched openShiftPipelinesConfig.
func (c *FakeOpenShiftPipelinesConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpenShiftPipelinesConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(openshiftpipelinesconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpenShiftPipelinesConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftPipelinesConfig), err
}