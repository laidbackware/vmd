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

	"github.com/laidbackware/vmd/api"
	"github.com/laidbackware/vmd/presenters"
	"github.com/spf13/cobra"
)

var version string

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, availability, err := api.ListFiles(slug, subProduct, version)
		if err != nil {
			handleErrors(err)
		}
		headings := []string{"Filename", "Size", "Build number", "Description"}

		presenters.PrintEntitlement(availability.EulaAccepted, availability.EligibleToDownload)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
