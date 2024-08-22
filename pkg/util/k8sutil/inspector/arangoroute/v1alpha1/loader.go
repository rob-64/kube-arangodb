//
// DISCLAIMER
//
// Copyright 2016-2024 ArangoDB GmbH, Cologne, Germany
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

package v1alpha1

import (
	networkingApi "github.com/arangodb/kube-arangodb/pkg/apis/networking/v1alpha1"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector/gvk"
)

type Inspector interface {
	gvk.GVK

	ListSimple() []*networkingApi.ArangoRoute
	GetSimple(name string) (*networkingApi.ArangoRoute, bool)
	Filter(filters ...Filter) []*networkingApi.ArangoRoute
	Iterate(action Action, filters ...Filter) error
	Read() ReadInterface
}

type Filter func(at *networkingApi.ArangoRoute) bool
type Action func(at *networkingApi.ArangoRoute) error

func FilterObject(at *networkingApi.ArangoRoute, filters ...Filter) bool {
	for _, f := range filters {
		if f == nil {
			continue
		}

		if !f(at) {
			return false
		}
	}

	return true
}