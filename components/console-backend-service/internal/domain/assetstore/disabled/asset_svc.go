// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import (
	v1alpha2 "github.com/kyma-project/kyma/components/asset-store-controller-manager/pkg/apis/assetstore/v1alpha2"
	resource "github.com/kyma-project/kyma/components/console-backend-service/pkg/resource"
)

// assetSvc is an autogenerated failing mock type for the assetSvc type
type assetSvc struct {
	err error
}

// NewAssetSvc creates a new assetSvc type instance
func NewAssetSvc(err error) *assetSvc {
	return &assetSvc{err: err}
}

// Find provides a failing mock function with given fields: namespace, name
func (_m *assetSvc) Find(namespace string, name string) (*v1alpha2.Asset, error) {
	var r0 *v1alpha2.Asset
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ListForDocsTopicByType provides a failing mock function with given fields: namespace, docsTopicName, types
func (_m *assetSvc) ListForDocsTopicByType(namespace string, docsTopicName string, types []string) ([]*v1alpha2.Asset, error) {
	var r0 []*v1alpha2.Asset
	var r1 error
	r1 = _m.err

	return r0, r1
}

// Subscribe provides a failing mock function with given fields: listener
func (_m *assetSvc) Subscribe(listener resource.Listener) {
}

// Unsubscribe provides a failing mock function with given fields: listener
func (_m *assetSvc) Unsubscribe(listener resource.Listener) {
}