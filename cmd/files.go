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

	"github.com/laidbackware/vmd/api"
	"github.com/laidbackware/vmd/presenters"
	"github.com/spf13/cobra"
)

var version string

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Aliases: []string{"f"},
	Short: "List available files",
	Long: `List available files of a version of a sub-product

Either VMD_USER and VMD_PASS environment variable must be set
or the --user and --pass flags should be added`,
	Example: getFiles,
	Run: func(cmd *cobra.Command, args []string) {
		validateCredentials(cmd)
		files, availability, err := api.ListFiles(slug, subProduct, version, username, password)
		if err != nil {
			handleErrors(err)
		}
		headings := []string{"Filename", "Size", "Build number", "Description"}

		fmt.Printf("\nEula Accepted:         %t\n", availability.EulaAccepted)
		fmt.Printf("Eligable to Download:  %t\n\n", availability.EligibleToDownload)
		presenters.RenderTable(headings, files)
	},
}

func init() {
	getCmd.AddCommand(filesCmd)
	filesCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	filesCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	filesCmd.Flags().StringVarP(&version, "version", "v", "", "Version string")
	filesCmd.MarkFlagRequired("product")
	filesCmd.MarkFlagRequired("sub-product")
	filesCmd.MarkFlagRequired("version")
}
