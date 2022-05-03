//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
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

package inspector

import (
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector/anonymous"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (p *arangoTasksInspector) Anonymous(gvk schema.GroupVersionKind) (anonymous.Interface, bool) {
	g := ArangoTaskGK()

	if g.Kind == gvk.Kind && g.Group == gvk.Group {
		switch gvk.Version {
		case ArangoTaskVersionV1, DefaultVersion:
			if p.v1 == nil || p.v1.err != nil {
				return nil, false
			}
			return &arangoTasksInspectorAnonymousV1{i: p.v1}, true
		}
	}

	return nil, false
}
