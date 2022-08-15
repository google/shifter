/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version  string = "development"
	Platform string = "platform"
)

func init() {
	//Version = "0.3.0 linux/amd64"
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version number of Shifter",
	Long:  `This is the version of Shifter you are running`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

----------------------------------------
https://github.com/google/shifter
`)
		fmt.Println("Shifter version " + Version + " " + Platform)
	},
}
