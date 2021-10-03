package api

import (
)

func GetEula(slug, subProduct, version, username, password string) (data string, err error) {
	var downloadGroup, productID string
	downloadGroup, productID, err = basicClient.GetDlgProduct(slug, subProduct, version)
	if err != nil {return}
	if err = EnsureLogin(username, password); err != nil {return}

	data, err = authenticatedClient.FetchEulaUrl(downloadGroup, productID)
	return
}