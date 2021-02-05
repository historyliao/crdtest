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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	versioned "github.com/historyliao/crdtest/lib/testresource/generated/clientset/versioned"
	internalinterfaces "github.com/historyliao/crdtest/lib/testresource/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/historyliao/crdtest/lib/testresource/generated/listers/testresource/v1beta1"
	testresourcev1beta1 "github.com/historyliao/crdtest/lib/testresource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TestResourceInformer provides access to a shared informer and lister for
// TestResources.
type TestResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.TestResourceLister
}

type testResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTestResourceInformer constructs a new informer for TestResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTestResourceInformer constructs a new informer for TestResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GithubV1beta1().TestResources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GithubV1beta1().TestResources(namespace).Watch(context.TODO(), options)
			},
		},
		&testresourcev1beta1.TestResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *testResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *testResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&testresourcev1beta1.TestResource{}, f.defaultInformer)
}

func (f *testResourceInformer) Lister() v1beta1.TestResourceLister {
	return v1beta1.NewTestResourceLister(f.Informer().GetIndexer())
}
