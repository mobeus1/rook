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
	v1alpha1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSWIFTs implements SWIFTInterface
type FakeSWIFTs struct {
	Fake *FakeEdgefsV1alpha1
	ns   string
}

var swiftsResource = schema.GroupVersionResource{Group: "edgefs.rook.io", Version: "v1alpha1", Resource: "swifts"}

var swiftsKind = schema.GroupVersionKind{Group: "edgefs.rook.io", Version: "v1alpha1", Kind: "SWIFT"}

// Get takes name of the sWIFT, and returns the corresponding sWIFT object, and an error if there is any.
func (c *FakeSWIFTs) Get(name string, options v1.GetOptions) (result *v1alpha1.SWIFT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(swiftsResource, c.ns, name), &v1alpha1.SWIFT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SWIFT), err
}

// List takes label and field selectors, and returns the list of SWIFTs that match those selectors.
func (c *FakeSWIFTs) List(opts v1.ListOptions) (result *v1alpha1.SWIFTList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(swiftsResource, swiftsKind, c.ns, opts), &v1alpha1.SWIFTList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SWIFTList{ListMeta: obj.(*v1alpha1.SWIFTList).ListMeta}
	for _, item := range obj.(*v1alpha1.SWIFTList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sWIFTs.
func (c *FakeSWIFTs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(swiftsResource, c.ns, opts))

}

// Create takes the representation of a sWIFT and creates it.  Returns the server's representation of the sWIFT, and an error, if there is any.
func (c *FakeSWIFTs) Create(sWIFT *v1alpha1.SWIFT) (result *v1alpha1.SWIFT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(swiftsResource, c.ns, sWIFT), &v1alpha1.SWIFT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SWIFT), err
}

// Update takes the representation of a sWIFT and updates it. Returns the server's representation of the sWIFT, and an error, if there is any.
func (c *FakeSWIFTs) Update(sWIFT *v1alpha1.SWIFT) (result *v1alpha1.SWIFT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(swiftsResource, c.ns, sWIFT), &v1alpha1.SWIFT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SWIFT), err
}

// Delete takes name of the sWIFT and deletes it. Returns an error if one occurs.
func (c *FakeSWIFTs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(swiftsResource, c.ns, name), &v1alpha1.SWIFT{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSWIFTs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(swiftsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SWIFTList{})
	return err
}

// Patch applies the patch and returns the patched sWIFT.
func (c *FakeSWIFTs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SWIFT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(swiftsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SWIFT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SWIFT), err
}
