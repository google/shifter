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
	"log"
	api "shifter/api"

	"github.com/spf13/cobra"
)

var (
	httpPort string
)

const (
	serverAddress = "0.0.0.0:8080"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Convert Openshift Resources to Kubernetes native formats via Shifter API",
	Long: `

	 _____ __    _ ______            
	/ ___// /_  (_) __/ /____  _____       ___   ____ ____ ______  ___    ___   ____
	\__ \/ __ \/ / /_/ __/ _ \/ ___/      / _ \ / __// __//_  __/ / _ |  / _ \ /  _/
   ___/ / / / / / __/ /_/  __/ /         / , _// _/ _\ \   / /   / __ | / ___/_/ /
  /____/_/ /_/_/_/  \__/\___/_/         /_/|_|/___//___/  /_/   /_/ |_|/_/   /___/ 
                                 

Convert OpenShift resources to kubernetes native formats

Usage: shifter server

`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(`
   _____ __    _ ______            
  / ___// /_  (_) __/ /____  _____       ___   ____ ____ ______  ___    ___   ____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/      / _ \ / __// __//_  __/ / _ |  / _ \ /  _/
 ___/ / / / / / __/ /_/  __/ /         / , _// _/ _\ \   / /   / __ | / ___/_/ /
/____/_/ /_/_/_/  \__/\___/_/         /_/|_|/___//___/  /_/   /_/ |_|/_/   /___/ 
                                 
-------------------------------------------------------------------------------------
			`)

		//flags := ProcFlags(pFlags)
		err := api.ServerStart(serverAddress)
		if err != nil {
			log.Fatal("Cannot Start HTTP Server:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&httpPort, "port", "p", "8080", "Server Port: Default 8080")
}
