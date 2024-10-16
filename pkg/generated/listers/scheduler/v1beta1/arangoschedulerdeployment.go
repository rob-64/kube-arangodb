//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/arangodb/kube-arangodb/pkg/apis/scheduler/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoSchedulerDeploymentLister helps list ArangoSchedulerDeployments.
// All objects returned here must be treated as read-only.
type ArangoSchedulerDeploymentLister interface {
	// List lists all ArangoSchedulerDeployments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ArangoSchedulerDeployment, err error)
	// ArangoSchedulerDeployments returns an object that can list and get ArangoSchedulerDeployments.
	ArangoSchedulerDeployments(namespace string) ArangoSchedulerDeploymentNamespaceLister
	ArangoSchedulerDeploymentListerExpansion
}

// arangoSchedulerDeploymentLister implements the ArangoSchedulerDeploymentLister interface.
type arangoSchedulerDeploymentLister struct {
	indexer cache.Indexer
}

// NewArangoSchedulerDeploymentLister returns a new ArangoSchedulerDeploymentLister.
func NewArangoSchedulerDeploymentLister(indexer cache.Indexer) ArangoSchedulerDeploymentLister {
	return &arangoSchedulerDeploymentLister{indexer: indexer}
}

// List lists all ArangoSchedulerDeployments in the indexer.
func (s *arangoSchedulerDeploymentLister) List(selector labels.Selector) (ret []*v1beta1.ArangoSchedulerDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ArangoSchedulerDeployment))
	})
	return ret, err
}

// ArangoSchedulerDeployments returns an object that can list and get ArangoSchedulerDeployments.
func (s *arangoSchedulerDeploymentLister) ArangoSchedulerDeployments(namespace string) ArangoSchedulerDeploymentNamespaceLister {
	return arangoSchedulerDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoSchedulerDeploymentNamespaceLister helps list and get ArangoSchedulerDeployments.
// All objects returned here must be treated as read-only.
type ArangoSchedulerDeploymentNamespaceLister interface {
	// List lists all ArangoSchedulerDeployments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ArangoSchedulerDeployment, err error)
	// Get retrieves the ArangoSchedulerDeployment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ArangoSchedulerDeployment, error)
	ArangoSchedulerDeploymentNamespaceListerExpansion
}

// arangoSchedulerDeploymentNamespaceLister implements the ArangoSchedulerDeploymentNamespaceLister
// interface.
type arangoSchedulerDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoSchedulerDeployments in the indexer for a given namespace.
func (s arangoSchedulerDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ArangoSchedulerDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ArangoSchedulerDeployment))
	})
	return ret, err
}

// Get retrieves the ArangoSchedulerDeployment from the indexer for a given namespace and name.
func (s arangoSchedulerDeploymentNamespaceLister) Get(name string) (*v1beta1.ArangoSchedulerDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("arangoschedulerdeployment"), name)
	}
	return obj.(*v1beta1.ArangoSchedulerDeployment), nil
}
