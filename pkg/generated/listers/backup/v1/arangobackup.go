//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
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

package v1

import (
	v1 "github.com/arangodb/kube-arangodb/pkg/apis/backup/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoBackupLister helps list ArangoBackups.
// All objects returned here must be treated as read-only.
type ArangoBackupLister interface {
	// List lists all ArangoBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoBackup, err error)
	// ArangoBackups returns an object that can list and get ArangoBackups.
	ArangoBackups(namespace string) ArangoBackupNamespaceLister
	ArangoBackupListerExpansion
}

// arangoBackupLister implements the ArangoBackupLister interface.
type arangoBackupLister struct {
	indexer cache.Indexer
}

// NewArangoBackupLister returns a new ArangoBackupLister.
func NewArangoBackupLister(indexer cache.Indexer) ArangoBackupLister {
	return &arangoBackupLister{indexer: indexer}
}

// List lists all ArangoBackups in the indexer.
func (s *arangoBackupLister) List(selector labels.Selector) (ret []*v1.ArangoBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoBackup))
	})
	return ret, err
}

// ArangoBackups returns an object that can list and get ArangoBackups.
func (s *arangoBackupLister) ArangoBackups(namespace string) ArangoBackupNamespaceLister {
	return arangoBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoBackupNamespaceLister helps list and get ArangoBackups.
// All objects returned here must be treated as read-only.
type ArangoBackupNamespaceLister interface {
	// List lists all ArangoBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoBackup, err error)
	// Get retrieves the ArangoBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ArangoBackup, error)
	ArangoBackupNamespaceListerExpansion
}

// arangoBackupNamespaceLister implements the ArangoBackupNamespaceLister
// interface.
type arangoBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoBackups in the indexer for a given namespace.
func (s arangoBackupNamespaceLister) List(selector labels.Selector) (ret []*v1.ArangoBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoBackup))
	})
	return ret, err
}

// Get retrieves the ArangoBackup from the indexer for a given namespace and name.
func (s arangoBackupNamespaceLister) Get(name string) (*v1.ArangoBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("arangobackup"), name)
	}
	return obj.(*v1.ArangoBackup), nil
}
