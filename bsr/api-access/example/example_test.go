// Copyright 2020-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFileDescriptorSet(t *testing.T) {
	t.Parallel()

	res, err := Example()

	require.NoError(t, err)

	fds := res.GetFileDescriptorSet()
	require.NotNil(t, fds)

	files := fds.GetFile()
	require.Len(t, files, 1)
	assert.Equal(t, "connectrpc/eliza/v1/eliza.proto", files[0].GetName())
}
