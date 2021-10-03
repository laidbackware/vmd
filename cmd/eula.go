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

// filesCmd represents the files command
var eulaCmd = &cobra.Command{
	Use:   "eula",
	Aliases: []string{"e"},
	Short: "Display the Eula of a product",
	Long: `Display the eula of a version of a sub-product

Either VMD_USER and VMD_PASS environment variable must be set
or the --user and --pass flags should be added`,
	Example: getFiles,
	Run: func(cmd *cobra.Command, args []string) {
		validateCredentials(cmd)
		eula, err := api.GetEula(slug, subProduct, version, username, password)
		handleErrors(err)
		fmt.Printf("Open the URL in your browser: %s\n", eula)
	},
}

func init() {
	getCmd.AddCommand(eulaCmd)
	eulaCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	eulaCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	eulaCmd.Flags().StringVarP(&version, "version", "v", "", "Version string")
	eulaCmd.MarkFlagRequired("product")
	eulaCmd.MarkFlagRequired("sub-product")
	eulaCmd.MarkFlagRequired("version")
}
