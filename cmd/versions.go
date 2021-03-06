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
	"github.com/spf13/cobra"
)

var subProduct string

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions",
	Aliases: []string{"v"},
	Short: "List available versions",
	Long: "List available versions of a sub-product",
	Example: getVersions,
	Run: func(cmd *cobra.Command, args []string) {
		versionString, err := api.ListVersions(slug, subProduct)
		if err != nil {
			handleErrors(err)
		}
		fmt.Println(versionString)
	},
}

func init() {
	getCmd.AddCommand(versionsCmd)
	versionsCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	versionsCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	versionsCmd.MarkFlagRequired("product")
	versionsCmd.MarkFlagRequired("sub-product")
}
