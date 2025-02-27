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
	v1alpha1 "Hybrid_Cloud/pkg/apis/resource/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HCPPodLister helps list HCPPods.
// All objects returned here must be treated as read-only.
type HCPPodLister interface {
	// List lists all HCPPods in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HCPPod, err error)
	// HCPPods returns an object that can list and get HCPPods.
	HCPPods(namespace string) HCPPodNamespaceLister
	HCPPodListerExpansion
}

// hCPPodLister implements the HCPPodLister interface.
type hCPPodLister struct {
	indexer cache.Indexer
}

// NewHCPPodLister returns a new HCPPodLister.
func NewHCPPodLister(indexer cache.Indexer) HCPPodLister {
	return &hCPPodLister{indexer: indexer}
}

// List lists all HCPPods in the indexer.
func (s *hCPPodLister) List(selector labels.Selector) (ret []*v1alpha1.HCPPod, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HCPPod))
	})
	return ret, err
}

// HCPPods returns an object that can list and get HCPPods.
func (s *hCPPodLister) HCPPods(namespace string) HCPPodNamespaceLister {
	return hCPPodNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HCPPodNamespaceLister helps list and get HCPPods.
// All objects returned here must be treated as read-only.
type HCPPodNamespaceLister interface {
	// List lists all HCPPods in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HCPPod, err error)
	// Get retrieves the HCPPod from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HCPPod, error)
	HCPPodNamespaceListerExpansion
}

// hCPPodNamespaceLister implements the HCPPodNamespaceLister
// interface.
type hCPPodNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HCPPods in the indexer for a given namespace.
func (s hCPPodNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HCPPod, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HCPPod))
	})
	return ret, err
}

// Get retrieves the HCPPod from the indexer for a given namespace and name.
func (s hCPPodNamespaceLister) Get(name string) (*v1alpha1.HCPPod, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hcppod"), name)
	}
	return obj.(*v1alpha1.HCPPod), nil
}
