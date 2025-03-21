// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterOverridePolicies implements ClusterOverridePolicyInterface
type FakeClusterOverridePolicies struct {
	Fake *FakeCoreV1alpha1
}

var clusteroverridepoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("clusteroverridepolicies")

var clusteroverridepoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterOverridePolicy")

// Get takes name of the clusterOverridePolicy, and returns the corresponding clusterOverridePolicy object, and an error if there is any.
func (c *FakeClusterOverridePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterOverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteroverridepoliciesResource, name), &v1alpha1.ClusterOverridePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOverridePolicy), err
}

// List takes label and field selectors, and returns the list of ClusterOverridePolicies that match those selectors.
func (c *FakeClusterOverridePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterOverridePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteroverridepoliciesResource, clusteroverridepoliciesKind, opts), &v1alpha1.ClusterOverridePolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterOverridePolicyList{ListMeta: obj.(*v1alpha1.ClusterOverridePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterOverridePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterOverridePolicies.
func (c *FakeClusterOverridePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteroverridepoliciesResource, opts))
}

// Create takes the representation of a clusterOverridePolicy and creates it.  Returns the server's representation of the clusterOverridePolicy, and an error, if there is any.
func (c *FakeClusterOverridePolicies) Create(ctx context.Context, clusterOverridePolicy *v1alpha1.ClusterOverridePolicy, opts v1.CreateOptions) (result *v1alpha1.ClusterOverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteroverridepoliciesResource, clusterOverridePolicy), &v1alpha1.ClusterOverridePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOverridePolicy), err
}

// Update takes the representation of a clusterOverridePolicy and updates it. Returns the server's representation of the clusterOverridePolicy, and an error, if there is any.
func (c *FakeClusterOverridePolicies) Update(ctx context.Context, clusterOverridePolicy *v1alpha1.ClusterOverridePolicy, opts v1.UpdateOptions) (result *v1alpha1.ClusterOverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteroverridepoliciesResource, clusterOverridePolicy), &v1alpha1.ClusterOverridePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOverridePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterOverridePolicies) UpdateStatus(ctx context.Context, clusterOverridePolicy *v1alpha1.ClusterOverridePolicy, opts v1.UpdateOptions) (*v1alpha1.ClusterOverridePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusteroverridepoliciesResource, "status", clusterOverridePolicy), &v1alpha1.ClusterOverridePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOverridePolicy), err
}

// Delete takes name of the clusterOverridePolicy and deletes it. Returns an error if one occurs.
func (c *FakeClusterOverridePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusteroverridepoliciesResource, name, opts), &v1alpha1.ClusterOverridePolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterOverridePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteroverridepoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterOverridePolicyList{})
	return err
}

// Patch applies the patch and returns the patched clusterOverridePolicy.
func (c *FakeClusterOverridePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterOverridePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteroverridepoliciesResource, name, pt, data, subresources...), &v1alpha1.ClusterOverridePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOverridePolicy), err
}
