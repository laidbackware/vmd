package api

import (
	"fmt"

	"github.com/laidbackware/vmware-download-sdk/sdk"
)

type Availability struct {
	EulaAccepted bool
	EligibleToDownload bool
}

func ListFiles(slug, subProduct, version string) (data [][]string, availability Availability, err error) {
	err = ensureLogin()
	if err != nil {return}
	var downloadGroup, productID string
	downloadGroup, productID, err = authenticatedClient.GetDlgProduct(slug, subProduct, version)
	if err != nil {return}

	fmt.Println("Getting DLG Details")
	var dlgDetails sdk.DlgDetails
	dlgDetails, err = authenticatedClient.GetDlgDetails(downloadGroup, productID)
	if err != nil {return}

	for _, v := range dlgDetails.DownloadDetails {
		if v.FileName != "" {
			line := []string{v.FileName, v.FileSize, v.Build, v.Title}
			data = append(data, line)
		}
	}

	availability = Availability{
		EulaAccepted: dlgDetails.EulaResponse.EulaAccepted,
		EligibleToDownload:  dlgDetails.EligibilityResponse.EligibleToDownload,
	}
	return
}

