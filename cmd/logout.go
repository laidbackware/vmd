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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove all session cookies",
	Long: "Remove all session cookies by deleting .vmd.cookies",
	Run: func(cmd *cobra.Command, args []string) {
		

		cookieFile := filepath.Join(homeDir(), ".vmd.cookies")
		if _, fileErr := os.Stat(cookieFile); os.IsNotExist(fileErr) {
			fmt.Println("No sessions cookies to delete")
			os.Exit(0)
		}
		err := os.Remove(cookieFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Unable to delete cookie file [%s].\n", cookieFile)
			fmt.Fprintf(os.Stderr, "%e", err)
			os.Exit(1)
		}
		fmt.Println("Deleted all session cookies")
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
