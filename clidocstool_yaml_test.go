// Copyright 2017 cli-docs-tool authors
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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:errcheck
func TestGenYamlTree(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "test-gen-yaml-tree")
	require.NoError(t, err)
	defer os.RemoveAll(tmpdir)

	c, err := New(Options{
		Root:      cmd,
		SourceDir: tmpdir,
		Plugin:    true,
	})
	require.NoError(t, err)
	require.NoError(t, c.GenYamlTree(cmd))

	fres := filepath.Join(tmpdir, "do_sub.yaml")
	require.FileExists(t, fres)
	bres, err := ioutil.ReadFile(fres)
	require.NoError(t, err)
	bexc, err := ioutil.ReadFile("fixtures/do_sub.yaml")
	require.NoError(t, err)
	assert.Equal(t, string(bres), string(bexc))
}
