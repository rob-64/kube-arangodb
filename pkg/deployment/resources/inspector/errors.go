//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
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

package inspector

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func newNotImplementedError(gvk schema.GroupVersionKind) error {
	return notImplementedError{gvk: gvk}
}

type notImplementedError struct {
	gvk schema.GroupVersionKind
}

func (n notImplementedError) Error() string {
	return fmt.Sprintf("Action is not implemented for %s", n.gvk.String())
}

func newInvalidTypeError(gvk schema.GroupVersionKind) error {
	return invalidTypeError{gvk: gvk}
}

type invalidTypeError struct {
	gvk schema.GroupVersionKind
}

func (n invalidTypeError) Error() string {
	return fmt.Sprintf("Type is invalid for %s", n.gvk.String())
}
