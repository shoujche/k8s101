/*
Copyright 2019 The Kubernetes Authors.

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
	time "time"

	demov1beta1 "github.com/jeremyxu2010/demo-apiserver/pkg/apis/demo/v1beta1"
	clientset "github.com/jeremyxu2010/demo-apiserver/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/jeremyxu2010/demo-apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1beta1 "github.com/jeremyxu2010/demo-apiserver/pkg/client/listers_generated/demo/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FooBarOptionsInformer provides access to a shared informer and lister for
// FooBarOptionses.
type FooBarOptionsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FooBarOptionsLister
}

type fooBarOptionsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFooBarOptionsInformer constructs a new informer for FooBarOptions type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFooBarOptionsInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFooBarOptionsInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFooBarOptionsInformer constructs a new informer for FooBarOptions type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFooBarOptionsInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DemoV1beta1().FooBarOptionses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DemoV1beta1().FooBarOptionses(namespace).Watch(options)
			},
		},
		&demov1beta1.FooBarOptions{},
		resyncPeriod,
		indexers,
	)
}

func (f *fooBarOptionsInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFooBarOptionsInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fooBarOptionsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&demov1beta1.FooBarOptions{}, f.defaultInformer)
}

func (f *fooBarOptionsInformer) Lister() v1beta1.FooBarOptionsLister {
	return v1beta1.NewFooBarOptionsLister(f.Informer().GetIndexer())
}
