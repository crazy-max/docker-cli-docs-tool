// Copyright 2021 cli-docs-tool authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clidocstool

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra/doc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:errcheck
func TestGenManTree(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "test-gen-man-tree")
	require.NoError(t, err)
	defer os.RemoveAll(tmpdir)

	c, err := New(Options{
		Root: cmd,
		ManHeader: &doc.GenManHeader{
			Title:   "DOCKER",
			Section: "1",
			Source:  "Docker Community",
			Manual:  "Docker User Manuals",
		},
		SourceDir: tmpdir,
		Plugin:    true,
	})
	require.NoError(t, err)
	require.NoError(t, c.GenManTree(cmd))

	fres := filepath.Join(tmpdir, "do-sub.1")
	require.FileExists(t, fres)
	bres, err := ioutil.ReadFile(fres)
	require.NoError(t, err)
	bexc, err := ioutil.ReadFile("fixtures/do-sub.1")
	require.NoError(t, err)
	assert.Equal(t, string(bres), string(bexc))
}
