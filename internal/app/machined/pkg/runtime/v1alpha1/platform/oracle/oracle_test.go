// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package oracle_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/siderolabs/talos/internal/app/machined/pkg/runtime/v1alpha1/platform/oracle"
)

//go:embed testdata/metadata.json
var rawMetadata []byte

//go:embed testdata/metadatanetwork.json
var rawMetadataNetwork []byte

//go:embed testdata/expected.yaml
var expectedNetworkConfig string

func TestParseMetadata(t *testing.T) {
	p := &oracle.Oracle{}

	var metadata oracle.MetadataConfig

	require.NoError(t, json.Unmarshal(rawMetadata, &metadata))

	var m []oracle.NetworkConfig

	require.NoError(t, json.Unmarshal(rawMetadataNetwork, &m))

	networkConfig, err := p.ParseMetadata(m, &metadata)
	require.NoError(t, err)

	marshaled, err := yaml.Marshal(networkConfig)
	require.NoError(t, err)

	assert.Equal(t, expectedNetworkConfig, string(marshaled))
}
