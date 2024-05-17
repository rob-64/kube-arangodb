//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/arangodb/kube-arangodb/pkg/apis/ml/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoMLExtensions implements ArangoMLExtensionInterface
type FakeArangoMLExtensions struct {
	Fake *FakeMlV1beta1
	ns   string
}

var arangomlextensionsResource = v1beta1.SchemeGroupVersion.WithResource("arangomlextensions")

var arangomlextensionsKind = v1beta1.SchemeGroupVersion.WithKind("ArangoMLExtension")

// Get takes name of the arangoMLExtension, and returns the corresponding arangoMLExtension object, and an error if there is any.
func (c *FakeArangoMLExtensions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ArangoMLExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(arangomlextensionsResource, c.ns, name), &v1beta1.ArangoMLExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArangoMLExtension), err
}

// List takes label and field selectors, and returns the list of ArangoMLExtensions that match those selectors.
func (c *FakeArangoMLExtensions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ArangoMLExtensionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(arangomlextensionsResource, arangomlextensionsKind, c.ns, opts), &v1beta1.ArangoMLExtensionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ArangoMLExtensionList{ListMeta: obj.(*v1beta1.ArangoMLExtensionList).ListMeta}
	for _, item := range obj.(*v1beta1.ArangoMLExtensionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoMLExtensions.
func (c *FakeArangoMLExtensions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(arangomlextensionsResource, c.ns, opts))

}

// Create takes the representation of a arangoMLExtension and creates it.  Returns the server's representation of the arangoMLExtension, and an error, if there is any.
func (c *FakeArangoMLExtensions) Create(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.CreateOptions) (result *v1beta1.ArangoMLExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(arangomlextensionsResource, c.ns, arangoMLExtension), &v1beta1.ArangoMLExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArangoMLExtension), err
}

// Update takes the representation of a arangoMLExtension and updates it. Returns the server's representation of the arangoMLExtension, and an error, if there is any.
func (c *FakeArangoMLExtensions) Update(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.UpdateOptions) (result *v1beta1.ArangoMLExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(arangomlextensionsResource, c.ns, arangoMLExtension), &v1beta1.ArangoMLExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArangoMLExtension), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoMLExtensions) UpdateStatus(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.UpdateOptions) (*v1beta1.ArangoMLExtension, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(arangomlextensionsResource, "status", c.ns, arangoMLExtension), &v1beta1.ArangoMLExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArangoMLExtension), err
}

// Delete takes name of the arangoMLExtension and deletes it. Returns an error if one occurs.
func (c *FakeArangoMLExtensions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(arangomlextensionsResource, c.ns, name, opts), &v1beta1.ArangoMLExtension{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoMLExtensions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(arangomlextensionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ArangoMLExtensionList{})
	return err
}

// Patch applies the patch and returns the patched arangoMLExtension.
func (c *FakeArangoMLExtensions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArangoMLExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(arangomlextensionsResource, c.ns, name, pt, data, subresources...), &v1beta1.ArangoMLExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArangoMLExtension), err
}