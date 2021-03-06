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

var slug string

// subproductsCmd represents the subproducts command
var subproductsCmd = &cobra.Command{
	Use:   "subproducts",
	Aliases: []string{"s"},
	Short: "List sub-products",
	Long: "List sub-products for a specified product",
	Example: getSubProductsUsage,
	Run: func(cmd *cobra.Command, args []string) {
		products, err := api.ListSubProducts(slug)
		if err != nil {
			handleErrors(err)
		}
		headings := []string{"Sub-Product Code", "Description"}
		presenters.RenderTable(headings, products)
	},
}

func init() {
	getCmd.AddCommand(subproductsCmd)
	subproductsCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	subproductsCmd.MarkFlagRequired("product")
}
