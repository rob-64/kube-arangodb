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

package s3

import "github.com/arangodb/kube-arangodb/pkg/util/errors"

type Config struct {
	Endpoint      string
	AllowInsecure bool
	CACrtFile     string
	CAKeyFile     string
	DisableSSL    bool
	Region        string
	BucketName    string
	AccessKeyFile string // path to file containing S3 AccessKey
	SecretKeyFile string // path to file containing S3 SecretKey
}

func (c Config) Validate() error {
	if c.AccessKeyFile == "" {
		return errors.Errorf("AccessKeyFile is not defined")
	}
	if c.SecretKeyFile == "" {
		return errors.Errorf("SecretKeyFile is not defined")
	}
	if c.Endpoint == "" {
		return errors.Errorf("Endpoint is not defined")
	}
	return nil
}
