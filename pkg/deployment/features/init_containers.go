//
// DISCLAIMER
//
// Copyright 2023-2024 ArangoDB GmbH, Cologne, Germany
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

package features

func init() {
	registerFeature(initContainerCopyResources)
	registerFeature(initContainerUpscaleResources)
}

var initContainerCopyResources = &feature{
	name:               "init-containers-copy-resources",
	description:        "Copy resources spec to built-in init containers if they are not specified",
	version:            "3.6.0",
	enterpriseRequired: false,
	enabledByDefault:   true,
	hidden:             false,
}

var initContainerUpscaleResources = &feature{
	name:               "init-containers-upscale-resources",
	description:        "Copy resources spec to built-in init containers if they are not specified or lower",
	version:            "3.6.0",
	enterpriseRequired: false,
	enabledByDefault:   true,
	hidden:             false,
}

func InitContainerCopyResources() Feature {
	return initContainerCopyResources
}

func InitContainerUpscaleResources() Feature {
	return initContainerUpscaleResources
}
