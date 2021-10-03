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
	// "fmt"

	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var getCmd = &cobra.Command{
	Use:   "get",
	Aliases: []string{"g"},
	Short: "Display responses",
	Long: `Display responses`,
	Example: fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s", getProductsUsage, getSubProductsUsage, getVersions, getFiles, getManifestExample),
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
