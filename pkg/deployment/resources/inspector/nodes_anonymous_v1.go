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
	"context"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type nodesInspectorAnonymousV1 struct {
	i *nodesInspectorV1
}

func (e *nodesInspectorAnonymousV1) Get(ctx context.Context, name string, opts meta.GetOptions) (meta.Object, error) {
	return e.i.Get(ctx, name, opts)
}

func (e *nodesInspectorAnonymousV1) Create(ctx context.Context, obj meta.Object, opts meta.CreateOptions) (meta.Object, error) {
	return nil, newNotImplementedError(NodeGKv1())
}

func (e *nodesInspectorAnonymousV1) Update(ctx context.Context, obj meta.Object, opts meta.UpdateOptions) (meta.Object, error) {
	return nil, newNotImplementedError(NodeGKv1())
}

func (e *nodesInspectorAnonymousV1) UpdateStatus(ctx context.Context, obj meta.Object, opts meta.UpdateOptions) (meta.Object, error) {
	return nil, newNotImplementedError(NodeGKv1())
}

func (e *nodesInspectorAnonymousV1) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts meta.PatchOptions, subresources ...string) (result meta.Object, err error) {
	return nil, newNotImplementedError(NodeGKv1())
}

func (e *nodesInspectorAnonymousV1) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
	return newNotImplementedError(NodeGKv1())
}
