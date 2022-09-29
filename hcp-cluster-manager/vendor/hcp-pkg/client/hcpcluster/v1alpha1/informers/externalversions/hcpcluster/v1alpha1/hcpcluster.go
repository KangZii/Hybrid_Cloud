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

package v1alpha1

import (
	hcpclusterv1alpha1 "hcp-pkg/apis/hcpcluster/v1alpha1"
	versioned "hcp-pkg/client/hcpcluster/v1alpha1/clientset/versioned"
	internalinterfaces "hcp-pkg/client/hcpcluster/v1alpha1/informers/externalversions/internalinterfaces"
	v1alpha1 "hcp-pkg/client/hcpcluster/v1alpha1/listers/hcpcluster/v1alpha1"
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HCPClusterInformer provides access to a shared informer and lister for
// HCPClusters.
type HCPClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HCPClusterLister
}

type hCPClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHCPClusterInformer constructs a new informer for HCPCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHCPClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHCPClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHCPClusterInformer constructs a new informer for HCPCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHCPClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HcpV1alpha1().HCPClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HcpV1alpha1().HCPClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&hcpclusterv1alpha1.HCPCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *hCPClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHCPClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hCPClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hcpclusterv1alpha1.HCPCluster{}, f.defaultInformer)
}

func (f *hCPClusterInformer) Lister() v1alpha1.HCPClusterLister {
	return v1alpha1.NewHCPClusterLister(f.Informer().GetIndexer())
}