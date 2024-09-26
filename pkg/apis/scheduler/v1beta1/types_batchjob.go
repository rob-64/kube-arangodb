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

package v1beta1

import (
	batch "k8s.io/api/batch/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/arangodb/kube-arangodb/pkg/apis/scheduler"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArangoSchedulerBatchJobList is a list of BatchJobs.
type ArangoSchedulerBatchJobList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ArangoSchedulerBatchJob `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArangoSchedulerBatchJob wraps batch. ArangoSchedulerBatchJob with profile details
type ArangoSchedulerBatchJob struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArangoSchedulerBatchJobSpec   `json:"spec"`
	Status ArangoSchedulerBatchJobStatus `json:"status"`
}

type ArangoSchedulerBatchJobSpec struct {
	// Profiles keeps list of the profiles
	Profiles []string `json:"profiles,omitempty"`

	batch.JobSpec `json:",inline"`
}

type ArangoSchedulerBatchJobStatus struct {
	ArangoSchedulerStatusMetadata `json:",inline"`

	batch.JobStatus `json:",inline"`
}

// AsOwner creates an OwnerReference for the given  ArangoSchedulerBatchJob
func (d *ArangoSchedulerBatchJob) AsOwner() meta.OwnerReference {
	trueVar := true
	return meta.OwnerReference{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       scheduler.BatchJobResourceKind,
		Name:       d.Name,
		UID:        d.UID,
		Controller: &trueVar,
	}
}

func (d *ArangoSchedulerBatchJob) GetStatus() ArangoSchedulerBatchJobStatus {
	return d.Status
}

func (d *ArangoSchedulerBatchJob) SetStatus(status ArangoSchedulerBatchJobStatus) {
	d.Status = status
}
