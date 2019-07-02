// Copyright 2017 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector_common

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/percona/mongodb_exporter/testutils"
)

func TestGetConnPoolStatsDecodesFine(t *testing.T) {
	// setup
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("mongod", func(t *testing.T) {
		// setup
		defaultClient := testutils.MustGetConnectedMongodClient(t, ctx)
		defer defaultClient.Disconnect(ctx)
		// run
		statusDefault := GetConnPoolStats(defaultClient)
		// test
		assert.NotNil(t, statusDefault)
	})

	t.Run("replset", func(t *testing.T) {
		// setup
		replSetClient := testutils.MustGetConnectedReplSetClient(t, ctx)
		defer replSetClient.Disconnect(ctx)
		// run
		statusReplSet := GetConnPoolStats(replSetClient)
		// test
		assert.NotNil(t, statusReplSet)
	})
}
