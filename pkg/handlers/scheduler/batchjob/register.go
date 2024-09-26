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

package batchjob

import (
	batch "k8s.io/api/batch/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	arangoClientSet "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	arangoInformer "github.com/arangodb/kube-arangodb/pkg/generated/informers/externalversions"
	"github.com/arangodb/kube-arangodb/pkg/handlers/generic/parent"
	operator "github.com/arangodb/kube-arangodb/pkg/operatorV2"
	"github.com/arangodb/kube-arangodb/pkg/operatorV2/event"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil"
)

// RegisterInformer into operator
func RegisterInformer(operator operator.Operator, recorder event.Recorder, client arangoClientSet.Interface,
	kubeClient kubernetes.Interface, informer arangoInformer.SharedInformerFactory, kubeInformer informers.SharedInformerFactory) error {

	if err := operator.RegisterInformer(informer.Scheduler().V1beta1().ArangoSchedulerBatchJobs().Informer(),
		Group(),
		Version(),
		Kind()); err != nil {
		return err
	}

	h := &handler{
		client:     client,
		kubeClient: kubeClient,

		eventRecorder: recorder.NewInstance(Group(), Version(), Kind()),

		operator: operator,
	}

	h.init()

	if err := operator.RegisterHandler(h); err != nil {
		return err
	}

	{
		batchJob := k8sutil.BatchV1JobGVK()

		if err := operator.RegisterInformer(kubeInformer.Batch().V1().Jobs().Informer(),
			batchJob.Group,
			batchJob.Version,
			batchJob.Kind); err != nil {
			return err
		}

		batchJobHandler := parent.NewNotifyHandler[*batch.Job]("batch-job-v1-parent", operator, kubeClient.BatchV1().Jobs, batchJob, GVK())

		if err := operator.RegisterHandler(batchJobHandler); err != nil {
			return err
		}
	}

	return nil
}
