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

package storage

import (
	"context"

	"github.com/arangodb/kube-arangodb/pkg/ml/storage/s3"
)

type StorageType string

const (
	StorageTypeS3Proxy = StorageType("s3")
)

type ServiceConfig struct {
	ListenAddress string
	S3            s3.Config
}

type Service interface {
	Run(ctx context.Context) error
}