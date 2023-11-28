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

package v1alpha1

import "github.com/arangodb/kube-arangodb/pkg/apis/shared"

type ArangoMLExtensionSpec struct {
	// MetadataService keeps the MetadataService configuration
	// +doc/immutable: This setting cannot be changed after the MetadataService has been created.
	MetadataService *ArangoMLExtensionSpecMetadataService `json:"metadataService,omitempty"`
}

func (a *ArangoMLExtensionSpec) GetMetadataService() *ArangoMLExtensionSpecMetadataService {
	if a == nil || a.MetadataService == nil {
		return nil
	}

	return a.MetadataService
}

func (a *ArangoMLExtensionSpec) Validate() error {
	return shared.WithErrors(shared.PrefixResourceErrors("spec",
		shared.PrefixResourceErrors("metadataService", a.GetMetadataService().Validate()),
	))
}
