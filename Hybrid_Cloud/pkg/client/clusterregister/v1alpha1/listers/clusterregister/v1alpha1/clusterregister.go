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
	v1alpha1 "Hybrid_Cloud/pkg/apis/clusterregister/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterRegisterLister helps list ClusterRegisters.
// All objects returned here must be treated as read-only.
type ClusterRegisterLister interface {
	// List lists all ClusterRegisters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterRegister, err error)
	// ClusterRegisters returns an object that can list and get ClusterRegisters.
	ClusterRegisters(namespace string) ClusterRegisterNamespaceLister
	ClusterRegisterListerExpansion
}

// clusterRegisterLister implements the ClusterRegisterLister interface.
type clusterRegisterLister struct {
	indexer cache.Indexer
}

// NewClusterRegisterLister returns a new ClusterRegisterLister.
func NewClusterRegisterLister(indexer cache.Indexer) ClusterRegisterLister {
	return &clusterRegisterLister{indexer: indexer}
}

// List lists all ClusterRegisters in the indexer.
func (s *clusterRegisterLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterRegister, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterRegister))
	})
	return ret, err
}

// ClusterRegisters returns an object that can list and get ClusterRegisters.
func (s *clusterRegisterLister) ClusterRegisters(namespace string) ClusterRegisterNamespaceLister {
	return clusterRegisterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterRegisterNamespaceLister helps list and get ClusterRegisters.
// All objects returned here must be treated as read-only.
type ClusterRegisterNamespaceLister interface {
	// List lists all ClusterRegisters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterRegister, err error)
	// Get retrieves the ClusterRegister from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterRegister, error)
	ClusterRegisterNamespaceListerExpansion
}

// clusterRegisterNamespaceLister implements the ClusterRegisterNamespaceLister
// interface.
type clusterRegisterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterRegisters in the indexer for a given namespace.
func (s clusterRegisterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterRegister, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterRegister))
	})
	return ret, err
}

// Get retrieves the ClusterRegister from the indexer for a given namespace and name.
func (s clusterRegisterNamespaceLister) Get(name string) (*v1alpha1.ClusterRegister, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterregister"), name)
	}
	return obj.(*v1alpha1.ClusterRegister), nil
}
