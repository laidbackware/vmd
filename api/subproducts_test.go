package api

import (
	"testing"

	"github.com/laidbackware/vmware-download-sdk/sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSubProducts(t *testing.T) {
	var products [][]string
	products, err := ListSubProducts("vmware_tools")
	require.Nil(t, err)
	assert.NotEmpty(t, products)
}

func TestGetSubProductsInvalidSlug(t *testing.T) {
	versions, err := ListVersions("tools", "vmtools")
	assert.ErrorIs(t, err, sdk.ErrorInvalidSlug)
	assert.Empty(t, versions, "Expected response to be empty")
}