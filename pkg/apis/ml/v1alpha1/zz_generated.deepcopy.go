//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	deploymentv1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	sharedv1 "github.com/arangodb/kube-arangodb/pkg/apis/shared/v1"
	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJob) DeepCopyInto(out *ArangoMLBatchJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJob.
func (in *ArangoMLBatchJob) DeepCopy() *ArangoMLBatchJob {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLBatchJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJobList) DeepCopyInto(out *ArangoMLBatchJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLBatchJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJobList.
func (in *ArangoMLBatchJobList) DeepCopy() *ArangoMLBatchJobList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLBatchJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJobSpec) DeepCopyInto(out *ArangoMLBatchJobSpec) {
	*out = *in
	if in.JobSpec != nil {
		in, out := &in.JobSpec, &out.JobSpec
		*out = new(v1.JobSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJobSpec.
func (in *ArangoMLBatchJobSpec) DeepCopy() *ArangoMLBatchJobSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLBatchJobStatus) DeepCopyInto(out *ArangoMLBatchJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(deploymentv1.ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.JobStatus != nil {
		in, out := &in.JobStatus, &out.JobStatus
		*out = new(v1.JobStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(sharedv1.Object)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLBatchJobStatus.
func (in *ArangoMLBatchJobStatus) DeepCopy() *ArangoMLBatchJobStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLBatchJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJob) DeepCopyInto(out *ArangoMLCronJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJob.
func (in *ArangoMLCronJob) DeepCopy() *ArangoMLCronJob {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLCronJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJobList) DeepCopyInto(out *ArangoMLCronJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLCronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJobList.
func (in *ArangoMLCronJobList) DeepCopy() *ArangoMLCronJobList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLCronJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJobSpec) DeepCopyInto(out *ArangoMLCronJobSpec) {
	*out = *in
	if in.CronJobSpec != nil {
		in, out := &in.CronJobSpec, &out.CronJobSpec
		*out = new(v1.CronJobSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJobSpec.
func (in *ArangoMLCronJobSpec) DeepCopy() *ArangoMLCronJobSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLCronJobStatus) DeepCopyInto(out *ArangoMLCronJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(deploymentv1.ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CronJobStatus != nil {
		in, out := &in.CronJobStatus, &out.CronJobStatus
		*out = new(v1.CronJobStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(sharedv1.Object)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLCronJobStatus.
func (in *ArangoMLCronJobStatus) DeepCopy() *ArangoMLCronJobStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLCronJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtension) DeepCopyInto(out *ArangoMLExtension) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtension.
func (in *ArangoMLExtension) DeepCopy() *ArangoMLExtension {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLExtension) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionList) DeepCopyInto(out *ArangoMLExtensionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionList.
func (in *ArangoMLExtensionList) DeepCopy() *ArangoMLExtensionList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLExtensionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpec) DeepCopyInto(out *ArangoMLExtensionSpec) {
	*out = *in
	if in.MetadataService != nil {
		in, out := &in.MetadataService, &out.MetadataService
		*out = new(ArangoMLExtensionSpecMetadataService)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(sharedv1.Object)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(sharedv1.Image)
		(*in).DeepCopyInto(*out)
	}
	if in.Init != nil {
		in, out := &in.Init, &out.Init
		*out = new(ArangoMLExtensionSpecInit)
		(*in).DeepCopyInto(*out)
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(ArangoMLExtensionSpecDeployment)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpec.
func (in *ArangoMLExtensionSpec) DeepCopy() *ArangoMLExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpecDeployment) DeepCopyInto(out *ArangoMLExtensionSpecDeployment) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ArangoMLExtensionSpecDeploymentService)
		(*in).DeepCopyInto(*out)
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(sharedv1.PodTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Prediction != nil {
		in, out := &in.Prediction, &out.Prediction
		*out = new(ArangoMLExtensionSpecDeploymentComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Training != nil {
		in, out := &in.Training, &out.Training
		*out = new(ArangoMLExtensionSpecDeploymentComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(ArangoMLExtensionSpecDeploymentComponent)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpecDeployment.
func (in *ArangoMLExtensionSpecDeployment) DeepCopy() *ArangoMLExtensionSpecDeployment {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpecDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpecDeploymentComponent) DeepCopyInto(out *ArangoMLExtensionSpecDeploymentComponent) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.ContainerTemplate != nil {
		in, out := &in.ContainerTemplate, &out.ContainerTemplate
		*out = new(sharedv1.ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpecDeploymentComponent.
func (in *ArangoMLExtensionSpecDeploymentComponent) DeepCopy() *ArangoMLExtensionSpecDeploymentComponent {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpecDeploymentComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpecDeploymentService) DeepCopyInto(out *ArangoMLExtensionSpecDeploymentService) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(corev1.ServiceType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpecDeploymentService.
func (in *ArangoMLExtensionSpecDeploymentService) DeepCopy() *ArangoMLExtensionSpecDeploymentService {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpecDeploymentService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpecInit) DeepCopyInto(out *ArangoMLExtensionSpecInit) {
	*out = *in
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(sharedv1.PodTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerTemplate != nil {
		in, out := &in.ContainerTemplate, &out.ContainerTemplate
		*out = new(sharedv1.ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpecInit.
func (in *ArangoMLExtensionSpecInit) DeepCopy() *ArangoMLExtensionSpecInit {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpecInit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpecMetadataService) DeepCopyInto(out *ArangoMLExtensionSpecMetadataService) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(ArangoMLExtensionSpecMetadataServiceLocal)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpecMetadataService.
func (in *ArangoMLExtensionSpecMetadataService) DeepCopy() *ArangoMLExtensionSpecMetadataService {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpecMetadataService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionSpecMetadataServiceLocal) DeepCopyInto(out *ArangoMLExtensionSpecMetadataServiceLocal) {
	*out = *in
	if in.ArangoPipeDatabase != nil {
		in, out := &in.ArangoPipeDatabase, &out.ArangoPipeDatabase
		*out = new(string)
		**out = **in
	}
	if in.ArangoMLFeatureStoreDatabase != nil {
		in, out := &in.ArangoMLFeatureStoreDatabase, &out.ArangoMLFeatureStoreDatabase
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionSpecMetadataServiceLocal.
func (in *ArangoMLExtensionSpecMetadataServiceLocal) DeepCopy() *ArangoMLExtensionSpecMetadataServiceLocal {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionSpecMetadataServiceLocal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionStatus) DeepCopyInto(out *ArangoMLExtensionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(deploymentv1.ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetadataService != nil {
		in, out := &in.MetadataService, &out.MetadataService
		*out = new(ArangoMLExtensionStatusMetadataService)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(sharedv1.ServiceAccount)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionStatus.
func (in *ArangoMLExtensionStatus) DeepCopy() *ArangoMLExtensionStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionStatusMetadataService) DeepCopyInto(out *ArangoMLExtensionStatusMetadataService) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(ArangoMLExtensionStatusMetadataServiceLocal)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(sharedv1.Object)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionStatusMetadataService.
func (in *ArangoMLExtensionStatusMetadataService) DeepCopy() *ArangoMLExtensionStatusMetadataService {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionStatusMetadataService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLExtensionStatusMetadataServiceLocal) DeepCopyInto(out *ArangoMLExtensionStatusMetadataServiceLocal) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLExtensionStatusMetadataServiceLocal.
func (in *ArangoMLExtensionStatusMetadataServiceLocal) DeepCopy() *ArangoMLExtensionStatusMetadataServiceLocal {
	if in == nil {
		return nil
	}
	out := new(ArangoMLExtensionStatusMetadataServiceLocal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorage) DeepCopyInto(out *ArangoMLStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorage.
func (in *ArangoMLStorage) DeepCopy() *ArangoMLStorage {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageList) DeepCopyInto(out *ArangoMLStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoMLStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageList.
func (in *ArangoMLStorageList) DeepCopy() *ArangoMLStorageList {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoMLStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageSpec) DeepCopyInto(out *ArangoMLStorageSpec) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.BucketPath != nil {
		in, out := &in.BucketPath, &out.BucketPath
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(ArangoMLStorageSpecMode)
		(*in).DeepCopyInto(*out)
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(ArangoMLStorageSpecBackend)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageSpec.
func (in *ArangoMLStorageSpec) DeepCopy() *ArangoMLStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageSpecBackend) DeepCopyInto(out *ArangoMLStorageSpecBackend) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(ArangoMLStorageSpecBackendS3)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageSpecBackend.
func (in *ArangoMLStorageSpecBackend) DeepCopy() *ArangoMLStorageSpecBackend {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageSpecBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageSpecBackendS3) DeepCopyInto(out *ArangoMLStorageSpecBackendS3) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(sharedv1.Object)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowInsecure != nil {
		in, out := &in.AllowInsecure, &out.AllowInsecure
		*out = new(bool)
		**out = **in
	}
	if in.CASecret != nil {
		in, out := &in.CASecret, &out.CASecret
		*out = new(sharedv1.Object)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageSpecBackendS3.
func (in *ArangoMLStorageSpecBackendS3) DeepCopy() *ArangoMLStorageSpecBackendS3 {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageSpecBackendS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageSpecMode) DeepCopyInto(out *ArangoMLStorageSpecMode) {
	*out = *in
	if in.Sidecar != nil {
		in, out := &in.Sidecar, &out.Sidecar
		*out = new(ArangoMLStorageSpecModeSidecar)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageSpecMode.
func (in *ArangoMLStorageSpecMode) DeepCopy() *ArangoMLStorageSpecMode {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageSpecMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageSpecModeSidecar) DeepCopyInto(out *ArangoMLStorageSpecModeSidecar) {
	*out = *in
	if in.ListenPort != nil {
		in, out := &in.ListenPort, &out.ListenPort
		*out = new(uint16)
		**out = **in
	}
	if in.ShutdownListenPort != nil {
		in, out := &in.ShutdownListenPort, &out.ShutdownListenPort
		*out = new(uint16)
		**out = **in
	}
	if in.ContainerTemplate != nil {
		in, out := &in.ContainerTemplate, &out.ContainerTemplate
		*out = new(sharedv1.ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageSpecModeSidecar.
func (in *ArangoMLStorageSpecModeSidecar) DeepCopy() *ArangoMLStorageSpecModeSidecar {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageSpecModeSidecar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoMLStorageStatus) DeepCopyInto(out *ArangoMLStorageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(deploymentv1.ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoMLStorageStatus.
func (in *ArangoMLStorageStatus) DeepCopy() *ArangoMLStorageStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoMLStorageStatus)
	in.DeepCopyInto(out)
	return out
}
