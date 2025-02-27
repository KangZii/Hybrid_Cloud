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

package v1alpha1

import (
	v1alpha1 "github.com/KETI-Hybrid/hcp-pkg/apis/hcpcluster/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HCPClusterLister helps list HCPClusters.
// All objects returned here must be treated as read-only.
type HCPClusterLister interface {
	// List lists all HCPClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HCPCluster, err error)
	// HCPClusters returns an object that can list and get HCPClusters.
	HCPClusters(namespace string) HCPClusterNamespaceLister
	HCPClusterListerExpansion
}

// hCPClusterLister implements the HCPClusterLister interface.
type hCPClusterLister struct {
	indexer cache.Indexer
}

// NewHCPClusterLister returns a new HCPClusterLister.
func NewHCPClusterLister(indexer cache.Indexer) HCPClusterLister {
	return &hCPClusterLister{indexer: indexer}
}

// List lists all HCPClusters in the indexer.
func (s *hCPClusterLister) List(selector labels.Selector) (ret []*v1alpha1.HCPCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HCPCluster))
	})
	return ret, err
}

// HCPClusters returns an object that can list and get HCPClusters.
func (s *hCPClusterLister) HCPClusters(namespace string) HCPClusterNamespaceLister {
	return hCPClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HCPClusterNamespaceLister helps list and get HCPClusters.
// All objects returned here must be treated as read-only.
type HCPClusterNamespaceLister interface {
	// List lists all HCPClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HCPCluster, err error)
	// Get retrieves the HCPCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HCPCluster, error)
	HCPClusterNamespaceListerExpansion
}

// hCPClusterNamespaceLister implements the HCPClusterNamespaceLister
// interface.
type hCPClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HCPClusters in the indexer for a given namespace.
func (s hCPClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HCPCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HCPCluster))
	})
	return ret, err
}

// Get retrieves the HCPCluster from the indexer for a given namespace and name.
func (s hCPClusterNamespaceLister) Get(name string) (*v1alpha1.HCPCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hcpcluster"), name)
	}
	return obj.(*v1alpha1.HCPCluster), nil
}
