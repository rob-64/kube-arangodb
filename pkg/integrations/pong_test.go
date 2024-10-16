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

package integrations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arangodb/kube-arangodb/pkg/util/shutdown"
)

func Test_Pong(t *testing.T) {
	c, health, internal, external := startService(t)
	defer c.Require(t)

	t.Run("Pong", func(t *testing.T) {
		t.Run("health", func(t *testing.T) {
			require.NoError(t, executeSync(t, shutdown.Context(),
				fmt.Sprintf("--address=127.0.0.1:%d", health),
				"--token=",
				"client",
				"pong",
				"v1"))
		})
		t.Run("internal", func(t *testing.T) {
			require.NoError(t, executeSync(t, shutdown.Context(),
				fmt.Sprintf("--address=127.0.0.1:%d", internal),
				"--token=",
				"client",
				"pong",
				"v1"))
		})
		t.Run("external", func(t *testing.T) {
			require.NoError(t, executeSync(t, shutdown.Context(),
				fmt.Sprintf("--address=127.0.0.1:%d", external),
				"--token=",
				"client",
				"pong",
				"v1"))
		})
	})
}
