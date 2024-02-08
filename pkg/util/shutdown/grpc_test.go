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

package shutdown

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/arangodb/kube-arangodb/pkg/api/server"
	pbShutdown "github.com/arangodb/kube-arangodb/pkg/api/shutdown/v1"
	"github.com/arangodb/kube-arangodb/pkg/util/svc"
	"github.com/arangodb/kube-arangodb/pkg/util/tests/tgrpc"
)

func Test_ShutdownGRPC(t *testing.T) {
	ctx, c := context.WithCancel(context.Background())
	defer c()

	local := svc.NewService(svc.Configuration{
		Address: "127.0.0.1:0",
	}, NewShutdownServer(c))

	start := local.Start(ctx)

	require.False(t, isContextDone(ctx))

	client := tgrpc.NewGRPCClient(t, ctx, pbShutdown.NewShutdownClient, start.Address())

	_, err := client.ShutdownServer(ctx, &server.Empty{})
	require.NoError(t, err)

	time.Sleep(time.Second)

	require.True(t, isContextDone(ctx))

	require.NoError(t, start.Wait())
}

func isContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
