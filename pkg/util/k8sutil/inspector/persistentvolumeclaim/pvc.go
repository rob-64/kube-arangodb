//
// DISCLAIMER
//
// Copyright 2021 ArangoDB GmbH, Cologne, Germany
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
// Author Adam Janikowski
//

package persistentvolumeclaim

import (
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector/refresh"
	core "k8s.io/api/core/v1"
)

type Inspector interface {
	refresh.Inspector

	PersistentVolumeClaims() []*core.PersistentVolumeClaim
	PersistentVolumeClaim(name string) (*core.PersistentVolumeClaim, bool)
	IteratePersistentVolumeClaims(action Action, filters ...Filter) error
	PersistentVolumeClaimReadInterface() ReadInterface
}

type Filter func(pvc *core.PersistentVolumeClaim) bool
type Action func(pvc *core.PersistentVolumeClaim) error
