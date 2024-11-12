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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/arangodb/kube-arangodb/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ArangoPlatformCharts returns a ArangoPlatformChartInformer.
	ArangoPlatformCharts() ArangoPlatformChartInformer
	// ArangoPlatformStorages returns a ArangoPlatformStorageInformer.
	ArangoPlatformStorages() ArangoPlatformStorageInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ArangoPlatformCharts returns a ArangoPlatformChartInformer.
func (v *version) ArangoPlatformCharts() ArangoPlatformChartInformer {
	return &arangoPlatformChartInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ArangoPlatformStorages returns a ArangoPlatformStorageInformer.
func (v *version) ArangoPlatformStorages() ArangoPlatformStorageInformer {
	return &arangoPlatformStorageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
