/*
Copyright © 2021 Matt Proud <proudmatt@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/laidbackware/vmd/api"
	"github.com/laidbackware/vmd/downloader"
	"github.com/laidbackware/vmd/manifest"
	"github.com/laidbackware/vmware-download-sdk/sdk"
	"github.com/spf13/cobra"
)

var(
	manifestFile string
	fileName string
	acceptEula bool
	outputDir string
)
	

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Aliases: []string{"d"},
	Short: "Download file from VMware",
	Long: `Download one or more files

Either VMD_USER and VMD_PASS environment variable must be set
or the --user and --pass flags should be added`,
	Example: downloadUsage,
	Run: func(cmd *cobra.Command, args []string) {
		validateCredentials(cmd)
		validateOutputDir()
		manifestWorkflow := validateDownloadFlags(cmd)
		err := api.EnsureLogin(username, password)
		handleErrors(err)
		if manifestWorkflow {
			downloadFromManifest()
		} else {
			fmt.Println("Collecting download payload")
			downloadPayloads, err := api.FetchDownloadPayload(slug, subProduct, version, fileName, username, password, acceptEula)
			handleErrors(err)
			downloadFiles(downloadPayloads)
		}
	},
}

func downloadFromManifest() {
	fmt.Printf("Opening manifest file: %s\n", manifestFile)
	manifestArray, err := manifest.ProcessFile(manifestFile)
	if err == manifest.ErrorFileDoesNotExist {
		fmt.Printf("File %s does not exist", manifestFile)
		os.Exit(1)
	} else if err == manifest.ErrorInvalidSpec {
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("Parsing file failed with error: %e\n", err)
		os.Exit(1)
	}

	var allPayloads [][]sdk.DownloadPayload
	for _, manifestSpec := range manifestArray {
		for _, glob := range manifestSpec.FilenameGlobs {
			fmt.Printf("Collecting download payload for [%s] [%s] [%s] [%s]\n", manifestSpec.Slug, manifestSpec.SubProduct, 
				manifestSpec.Version, glob)
			downloadPayloads, err := api.FetchDownloadPayload(manifestSpec.Slug, manifestSpec.SubProduct, manifestSpec.Version, 
				glob, username, password, acceptEula)
			handleErrors(err)
			allPayloads = append(allPayloads, downloadPayloads)
		}
	}

	for _, downloadPayloads := range allPayloads {
		downloadFiles(downloadPayloads)
	}

}

func downloadFiles(downloadPayloads []sdk.DownloadPayload) {
		for _, downloadPayload := range downloadPayloads {
			authorizedDownload, err := api.FetchDownloadLink(downloadPayload, username, password)
			handleErrors(err)
			authorizedDownload.FileName = filepath.Join(outputDir, authorizedDownload.FileName)
			err = downloader.TriggerDownload(authorizedDownload)
			handleErrors(err)
		}
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	downloadCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	downloadCmd.Flags().StringVarP(&version, "version", "v", "", "Version string. Can contain a glob.")
	downloadCmd.Flags().StringVarP(&fileName, "filename", "f", "", "Filename string. Can contain one or more globs.")
	downloadCmd.Flags().StringVarP(&manifestFile, "manifest", "m", "", "Filename of the manifest containing details of what to download")
	downloadCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Directory to download files to")
	downloadCmd.Flags().BoolVarP(&acceptEula, "accepteula", "a", false, "Filename string")
}
