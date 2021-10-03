package api

import (
	"testing"

	"github.com/laidbackware/vmware-download-sdk/sdk"
	"github.com/stretchr/testify/assert"
)

func TestGetEula(t *testing.T) {
	eulaUrl, err := GetEula("vmware_tools", "vmtools", "11.1.1", testing_user, testing_pass)
	assert.Nil(t, err)
	assert.NotEmpty(t, eulaUrl)
}

func TestGetEulaInvalidSlug(t *testing.T) {
	eulaUrl, err := GetEula("tools", "vmtools", "", testing_user, testing_pass)
	assert.ErrorIs(t, err, sdk.ErrorInvalidSlug)
	assert.Empty(t, eulaUrl)
}

func TestGetEulaInvalidSubProduct(t *testing.T) {
	eulaUrl, err := GetEula("vmware_tools", "tools", "", testing_user, testing_pass)
	assert.ErrorIs(t, err, sdk.ErrorInvalidSubProduct)
	assert.Empty(t, eulaUrl)
}

func TestGetEulaInvalidVersion(t *testing.T) {
	eulaUrl, err := GetEula("vmware_tools", "vmtools", "666", testing_user, testing_pass)
	assert.ErrorIs(t, err, sdk.ErrorInvalidVersion)
	assert.Empty(t, eulaUrl)
}
