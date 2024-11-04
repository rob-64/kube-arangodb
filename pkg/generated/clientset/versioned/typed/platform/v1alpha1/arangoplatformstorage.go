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

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/platform/v1alpha1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ArangoPlatformStoragesGetter has a method to return a ArangoPlatformStorageInterface.
// A group's client should implement this interface.
type ArangoPlatformStoragesGetter interface {
	ArangoPlatformStorages(namespace string) ArangoPlatformStorageInterface
}

// ArangoPlatformStorageInterface has methods to work with ArangoPlatformStorage resources.
type ArangoPlatformStorageInterface interface {
	Create(ctx context.Context, arangoPlatformStorage *v1alpha1.ArangoPlatformStorage, opts v1.CreateOptions) (*v1alpha1.ArangoPlatformStorage, error)
	Update(ctx context.Context, arangoPlatformStorage *v1alpha1.ArangoPlatformStorage, opts v1.UpdateOptions) (*v1alpha1.ArangoPlatformStorage, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, arangoPlatformStorage *v1alpha1.ArangoPlatformStorage, opts v1.UpdateOptions) (*v1alpha1.ArangoPlatformStorage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ArangoPlatformStorage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ArangoPlatformStorageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ArangoPlatformStorage, err error)
	ArangoPlatformStorageExpansion
}

// arangoPlatformStorages implements ArangoPlatformStorageInterface
type arangoPlatformStorages struct {
	*gentype.ClientWithList[*v1alpha1.ArangoPlatformStorage, *v1alpha1.ArangoPlatformStorageList]
}

// newArangoPlatformStorages returns a ArangoPlatformStorages
func newArangoPlatformStorages(c *PlatformV1alpha1Client, namespace string) *arangoPlatformStorages {
	return &arangoPlatformStorages{
		gentype.NewClientWithList[*v1alpha1.ArangoPlatformStorage, *v1alpha1.ArangoPlatformStorageList](
			"arangoplatformstorages",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.ArangoPlatformStorage { return &v1alpha1.ArangoPlatformStorage{} },
			func() *v1alpha1.ArangoPlatformStorageList { return &v1alpha1.ArangoPlatformStorageList{} }),
	}
}
