package api

import (
	"testing"

	"github.com/laidbackware/vmware-download-sdk/sdk"
	"github.com/stretchr/testify/assert"
)

func TestGetFiles(t *testing.T) {
	files, availability, err := ListFiles("vmware_tools", "vmtools", "11.1.1")
	assert.Nil(t, err)
	assert.Greater(t, len(files), 5, "Expected response to contain at least 5 items")
	assert.True(t, availability.EligibleToDownload)
}

func TestGetFilesInvalidSlug(t *testing.T) {
	files, _, err := ListFiles("tools", "vmtools", "")
	assert.ErrorIs(t, err, sdk.ErrorInvalidSlug)
	assert.Empty(t, files, "Expected response to be empty")
}

func TestGetFilesInvalidSubProduct(t *testing.T) {
	files, _, err := ListFiles("vmware_tools", "tools", "")
	assert.ErrorIs(t, err, sdk.ErrorInvalidSubProduct)
	assert.Empty(t, files, "Expected response to be empty")
}

func TestGetFilesInvalidVersion(t *testing.T) {
	files, _, err := ListFiles("vmware_tools", "vmtools", "666")
	assert.ErrorIs(t, err, sdk.ErrorInvalidVersion)
	assert.Empty(t, files, "Expected response to be empty")
}

func TestGetFilesNotEntitled(t *testing.T) {
	files, availability, err := ListFiles("vmware_nsx_t_data_center", "nsx-t", "3.1.3.1")
	assert.Nil(t, err)
	assert.Greater(t, len(files), 5, "Expected response to contain at least 5 items")
	assert.False(t, availability.EligibleToDownload)
}
