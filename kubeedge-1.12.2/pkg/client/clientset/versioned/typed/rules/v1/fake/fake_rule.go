/*
Copyright The KubeEdge Authors.

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

	rulesv1 "github.com/kubeedge/kubeedge/pkg/apis/rules/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRules implements RuleInterface
type FakeRules struct {
	Fake *FakeRulesV1
	ns   string
}

var rulesResource = schema.GroupVersionResource{Group: "rules.kubeedge.io", Version: "v1", Resource: "rules"}

var rulesKind = schema.GroupVersionKind{Group: "rules.kubeedge.io", Version: "v1", Kind: "Rule"}

// Get takes name of the rule, and returns the corresponding rule object, and an error if there is any.
func (c *FakeRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *rulesv1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rulesResource, c.ns, name), &rulesv1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.Rule), err
}

// List takes label and field selectors, and returns the list of Rules that match those selectors.
func (c *FakeRules) List(ctx context.Context, opts v1.ListOptions) (result *rulesv1.RuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rulesResource, rulesKind, c.ns, opts), &rulesv1.RuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rulesv1.RuleList{ListMeta: obj.(*rulesv1.RuleList).ListMeta}
	for _, item := range obj.(*rulesv1.RuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rules.
func (c *FakeRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rulesResource, c.ns, opts))

}

// Create takes the representation of a rule and creates it.  Returns the server's representation of the rule, and an error, if there is any.
func (c *FakeRules) Create(ctx context.Context, rule *rulesv1.Rule, opts v1.CreateOptions) (result *rulesv1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rulesResource, c.ns, rule), &rulesv1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.Rule), err
}

// Update takes the representation of a rule and updates it. Returns the server's representation of the rule, and an error, if there is any.
func (c *FakeRules) Update(ctx context.Context, rule *rulesv1.Rule, opts v1.UpdateOptions) (result *rulesv1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rulesResource, c.ns, rule), &rulesv1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.Rule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRules) UpdateStatus(ctx context.Context, rule *rulesv1.Rule, opts v1.UpdateOptions) (*rulesv1.Rule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rulesResource, "status", c.ns, rule), &rulesv1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.Rule), err
}

// Delete takes name of the rule and deletes it. Returns an error if one occurs.
func (c *FakeRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rulesResource, c.ns, name, opts), &rulesv1.Rule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &rulesv1.RuleList{})
	return err
}

// Patch applies the patch and returns the patched rule.
func (c *FakeRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *rulesv1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rulesResource, c.ns, name, pt, data, subresources...), &rulesv1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.Rule), err
}
