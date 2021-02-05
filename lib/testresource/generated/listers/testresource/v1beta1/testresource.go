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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/historyliao/crdtest/lib/testresource/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TestResourceLister helps list TestResources.
// All objects returned here must be treated as read-only.
type TestResourceLister interface {
	// List lists all TestResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.TestResource, err error)
	// TestResources returns an object that can list and get TestResources.
	TestResources(namespace string) TestResourceNamespaceLister
	TestResourceListerExpansion
}

// testResourceLister implements the TestResourceLister interface.
type testResourceLister struct {
	indexer cache.Indexer
}

// NewTestResourceLister returns a new TestResourceLister.
func NewTestResourceLister(indexer cache.Indexer) TestResourceLister {
	return &testResourceLister{indexer: indexer}
}

// List lists all TestResources in the indexer.
func (s *testResourceLister) List(selector labels.Selector) (ret []*v1beta1.TestResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TestResource))
	})
	return ret, err
}

// TestResources returns an object that can list and get TestResources.
func (s *testResourceLister) TestResources(namespace string) TestResourceNamespaceLister {
	return testResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TestResourceNamespaceLister helps list and get TestResources.
// All objects returned here must be treated as read-only.
type TestResourceNamespaceLister interface {
	// List lists all TestResources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.TestResource, err error)
	// Get retrieves the TestResource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.TestResource, error)
	TestResourceNamespaceListerExpansion
}

// testResourceNamespaceLister implements the TestResourceNamespaceLister
// interface.
type testResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TestResources in the indexer for a given namespace.
func (s testResourceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.TestResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TestResource))
	})
	return ret, err
}

// Get retrieves the TestResource from the indexer for a given namespace and name.
func (s testResourceNamespaceLister) Get(name string) (*v1beta1.TestResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("testresource"), name)
	}
	return obj.(*v1beta1.TestResource), nil
}
