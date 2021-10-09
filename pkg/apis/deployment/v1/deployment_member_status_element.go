//
// DISCLAIMER
//
// Copyright 2016-2021 ArangoDB GmbH, Cologne, Germany
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

package v1

type DeploymentStatusMemberElements []DeploymentStatusMemberElement

// DeploymentStatusMemberElement holds one specific element with group and member status
type DeploymentStatusMemberElement struct {
	Group  ServerGroup  `json:"group,omitempty"`
	Member MemberStatus `json:"member,omitempty"`
}

func (ds DeploymentStatusMembers) AsList() DeploymentStatusMemberElements {
	return ds.AsListInGroups(AllServerGroups...)
}

func (ds DeploymentStatusMembers) AsListInGroups(groups ...ServerGroup) DeploymentStatusMemberElements {
	var elements []DeploymentStatusMemberElement

	// Always return nil, so no error handling
	for _, g := range groups {
		elements = append(elements, ds.AsListInGroup(g)...)
	}

	return elements
}

func (ds DeploymentStatusMembers) AsListInGroup(group ServerGroup) DeploymentStatusMemberElements {
	var r DeploymentStatusMemberElements

	for _, m := range ds.MembersOfGroup(group) {
		r = append(r, DeploymentStatusMemberElement{
			Group:  group,
			Member: m,
		})
	}

	return r
}