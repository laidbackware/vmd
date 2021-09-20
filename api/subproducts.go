package api

import (
	// "fmt"
	// "os"

	"github.com/laidbackware/vmware-download-sdk/sdk"
)

func ListSubProducts(slug string) (data [][]string, err error) {
	var subProducts []sdk.SubProduct
	subProducts, err = basicClient.GetSubProductsSlice(slug)
	if err != nil {return}
	for _, v := range subProducts {
		line := []string{v.ProductCode, v.ProductName}
		data = append(data, line)
	}

	return
}

