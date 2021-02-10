//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
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

package pod

import (
	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"github.com/arangodb/kube-arangodb/pkg/deployment/features"
	"github.com/arangodb/kube-arangodb/pkg/deployment/resources/inspector"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil"
	core "k8s.io/api/core/v1"
)

func UpgradeVersionCheck() Builder {
	return upgradeVersionCheck{}
}

type upgradeVersionCheck struct{}

func (u upgradeVersionCheck) Args(i Input) k8sutil.OptionPairs {
	if features.UpgradeVersionCheck().Enabled() {
		switch i.Group {
		case api.ServerGroupAgents, api.ServerGroupDBServers, api.ServerGroupSingle:
			return k8sutil.NewOptionPair(k8sutil.OptionPair{
				Key:   "--database.check-version",
				Value: "true",
			})
		}
	}

	return nil
}

func (u upgradeVersionCheck) Volumes(i Input) ([]core.Volume, []core.VolumeMount) {
	return nil, nil
}

func (u upgradeVersionCheck) Envs(i Input) []core.EnvVar {
	return nil
}

func (u upgradeVersionCheck) Verify(i Input, cachedStatus inspector.Inspector) error {
	return nil
}
