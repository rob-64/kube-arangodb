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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "github.com/arangodb/kube-arangodb/pkg/apis/ml/v1beta1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ArangoMLStoragesGetter has a method to return a ArangoMLStorageInterface.
// A group's client should implement this interface.
type ArangoMLStoragesGetter interface {
	ArangoMLStorages(namespace string) ArangoMLStorageInterface
}

// ArangoMLStorageInterface has methods to work with ArangoMLStorage resources.
type ArangoMLStorageInterface interface {
	Create(ctx context.Context, arangoMLStorage *v1beta1.ArangoMLStorage, opts v1.CreateOptions) (*v1beta1.ArangoMLStorage, error)
	Update(ctx context.Context, arangoMLStorage *v1beta1.ArangoMLStorage, opts v1.UpdateOptions) (*v1beta1.ArangoMLStorage, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, arangoMLStorage *v1beta1.ArangoMLStorage, opts v1.UpdateOptions) (*v1beta1.ArangoMLStorage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ArangoMLStorage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ArangoMLStorageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArangoMLStorage, err error)
	ArangoMLStorageExpansion
}

// arangoMLStorages implements ArangoMLStorageInterface
type arangoMLStorages struct {
	*gentype.ClientWithList[*v1beta1.ArangoMLStorage, *v1beta1.ArangoMLStorageList]
}

// newArangoMLStorages returns a ArangoMLStorages
func newArangoMLStorages(c *MlV1beta1Client, namespace string) *arangoMLStorages {
	return &arangoMLStorages{
		gentype.NewClientWithList[*v1beta1.ArangoMLStorage, *v1beta1.ArangoMLStorageList](
			"arangomlstorages",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.ArangoMLStorage { return &v1beta1.ArangoMLStorage{} },
			func() *v1beta1.ArangoMLStorageList { return &v1beta1.ArangoMLStorageList{} }),
	}
}
