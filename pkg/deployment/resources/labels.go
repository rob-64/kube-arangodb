//
// DISCLAIMER
//
// Copyright 2020-2021 ArangoDB GmbH, Cologne, Germany
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
// Author Tomasz Mielech
//

package resources

import (
	"context"

	"github.com/arangodb/kube-arangodb/pkg/util/errors"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil"
	inspectorInterface "github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector"
	monitoring "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	core "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (r *Resources) EnsureLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	r.log.Info().Msgf("Ensuring labels")

	if err := r.EnsureSecretLabels(ctx, cachedStatus); err != nil {
		return err
	}

	if err := r.EnsureServiceAccountsLabels(ctx, cachedStatus); err != nil {
		return err
	}

	if err := r.EnsureServicesLabels(ctx, cachedStatus); err != nil {
		return err
	}

	if err := r.EnsureServiceMonitorsLabels(ctx, cachedStatus); err != nil {
		return err
	}

	if err := r.EnsurePodsLabels(ctx, cachedStatus); err != nil {
		return err
	}

	if err := r.EnsurePersistentVolumeClaimsLabels(ctx, cachedStatus); err != nil {
		return err
	}

	if err := r.EnsurePodDisruptionBudgetsLabels(ctx, cachedStatus); err != nil {
		return err
	}

	return nil
}

func (r *Resources) EnsureSecretLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	changed := false
	if err := cachedStatus.IterateSecrets(func(secret *core.Secret) error {
		if ensureLabelsMap(secret.Kind, secret, r.context.GetSpec(), func(name string, d []byte) error {
			return k8sutil.RunWithTimeout(ctx, func(ctxChild context.Context) error {
				_, err := r.context.SecretsModInterface().Patch(ctxChild,
					name, types.JSONPatchType, d, meta.PatchOptions{})
				return err
			})
		}) {
			changed = true
		}

		return nil
	}, func(secret *core.Secret) bool {
		return r.isChildResource(secret)
	}); err != nil {
		return err
	}

	if changed {
		return errors.Reconcile()
	}

	return nil
}

func (r *Resources) EnsureServiceAccountsLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	changed := false
	if err := cachedStatus.IterateServiceAccounts(func(serviceAccount *core.ServiceAccount) error {
		if ensureLabelsMap(serviceAccount.Kind, serviceAccount, r.context.GetSpec(), func(name string, d []byte) error {
			return k8sutil.RunWithTimeout(ctx, func(ctxChild context.Context) error {
				_, err := r.context.ServiceAccountsModInterface().Patch(ctxChild, name, types.JSONPatchType, d, meta.PatchOptions{})
				return err
			})
		}) {
			changed = true
		}

		return nil
	}, func(serviceAccount *core.ServiceAccount) bool {
		return r.isChildResource(serviceAccount)
	}); err != nil {
		return err
	}

	if changed {
		return errors.Reconcile()
	}

	return nil
}

func (r *Resources) EnsureServicesLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	changed := false
	if err := cachedStatus.IterateServices(func(service *core.Service) error {
		if ensureLabelsMap(service.Kind, service, r.context.GetSpec(), func(name string, d []byte) error {
			return k8sutil.RunWithTimeout(ctx, func(ctxChild context.Context) error {
				_, err := r.context.ServicesModInterface().Patch(ctxChild, name, types.JSONPatchType, d, meta.PatchOptions{})
				return err
			})
		}) {
			changed = true
		}

		return nil
	}, func(service *core.Service) bool {
		return r.isChildResource(service)
	}); err != nil {
		return err
	}

	if changed {
		return errors.Reconcile()
	}

	return nil
}

func (r *Resources) EnsureServiceMonitorsLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	changed := false
	if err := cachedStatus.IterateServiceMonitors(func(serviceMonitor *monitoring.ServiceMonitor) error {
		if ensureLabelsMap(serviceMonitor.Kind, serviceMonitor, r.context.GetSpec(), func(name string, d []byte) error {
			return k8sutil.RunWithTimeout(ctx, func(ctxChild context.Context) error {
				_, err := r.context.ServiceMonitorsModInterface().Patch(ctxChild, name, types.JSONPatchType, d, meta.PatchOptions{})
				return err
			})
		}) {
			changed = true
		}

		return nil
	}, func(serviceMonitor *monitoring.ServiceMonitor) bool {
		return r.isChildResource(serviceMonitor)
	}); err != nil {
		return err
	}

	if changed {
		return errors.Reconcile()
	}

	return nil
}

func (r *Resources) EnsurePodsLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	changed := false
	if err := cachedStatus.IteratePods(func(pod *core.Pod) error {
		if ensureGroupLabelsMap(pod.Kind, pod, r.context.GetSpec(), func(name string, d []byte) error {
			return k8sutil.RunWithTimeout(ctx, func(ctxChild context.Context) error {
				_, err := r.context.PodsModInterface().Patch(ctxChild, name, types.JSONPatchType, d, meta.PatchOptions{})
				return err
			})
		}) {
			changed = true
		}

		return nil
	}, func(pod *core.Pod) bool {
		return r.isChildResource(pod)
	}); err != nil {
		return err
	}

	if changed {
		return errors.Reconcile()
	}

	return nil
}

func (r *Resources) EnsurePersistentVolumeClaimsLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	changed := false
	if err := cachedStatus.IteratePersistentVolumeClaims(func(persistentVolumeClaim *core.PersistentVolumeClaim) error {
		if ensureGroupLabelsMap(persistentVolumeClaim.Kind, persistentVolumeClaim, r.context.GetSpec(), func(name string, d []byte) error {
			return k8sutil.RunWithTimeout(ctx, func(ctxChild context.Context) error {
				_, err := r.context.PersistentVolumeClaimsModInterface().Patch(ctxChild, name, types.JSONPatchType, d, meta.PatchOptions{})
				return err
			})
		}) {
			changed = true
		}

		return nil
	}, func(persistentVolumeClaim *core.PersistentVolumeClaim) bool {
		return r.isChildResource(persistentVolumeClaim)
	}); err != nil {
		return err
	}

	if changed {
		return errors.Reconcile()
	}

	return nil
}

func (r *Resources) EnsurePodDisruptionBudgetsLabels(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	changed := false
	if err := cachedStatus.IteratePodDisruptionBudgets(func(budget *policy.PodDisruptionBudget) error {
		if ensureLabelsMap(budget.Kind, budget, r.context.GetSpec(), func(name string, d []byte) error {
			return k8sutil.RunWithTimeout(ctx, func(ctxChild context.Context) error {
				_, err := r.context.PodDisruptionBudgetsModInterface().Patch(ctxChild, name, types.JSONPatchType, d, meta.PatchOptions{})
				return err
			})
		}) {
			changed = true
		}

		return nil
	}, func(budget *policy.PodDisruptionBudget) bool {
		return r.isChildResource(budget)
	}); err != nil {
		return err
	}

	if changed {
		return errors.Reconcile()
	}

	return nil
}
