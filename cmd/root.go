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

	"github.com/spf13/cobra"
)

var username string
var password string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.1.0",
	Use:   "vmd",
	Short: "Download binaries from customerconnect.vmware.com",
	Long: "vmd downloads binaries from customerconnect.vmware.com",
	Example: fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s", downloadUsage, getProductsUsage, getSubProductsUsage, getVersions, getFiles, getManifestExample),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&username, "user", "", "Username used to authenticate [$VMD_USER]")
	rootCmd.PersistentFlags().StringVar(&password, "pass", "", "Password used to authenticate [$VMD_PASS]")
}