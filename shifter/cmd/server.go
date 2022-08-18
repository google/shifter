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
	serverPort    string
	serverAddress string
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

		// Welcome Banners
		log.Printf("👋 INFO: Welcome to Shifter Server")
		log.Printf("🎬 INFO: Let's Start Shifting...")

		// Instanciate Shifter Server Instance
		server, err := api.InitServer(serverAddress, serverPort, sourcePath, outputPath)
		if err != nil {
			// Unable to instanciate Shifter HTTP Server
			log.Fatal("🌐 ❌ ERROR: Cannot Create Shifter HTTP Server:", err)
		}
		// Start Shifter Server Instance
		err = server.Start()
		if err != nil {
			// Unable to start Shifter HTTP Server
			log.Fatal("🌐 ❌ ERROR: Cannot Start Shiter HTTP Server:", err)
		}
	},
}

func init() {
	// TODO - Revisit the Flags and Required Flags, Fix descriptions add options, Add valdations.
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "8080", "Server Port: Default 8080")
	serverCmd.Flags().StringVarP(&serverAddress, "host-address", "a", "0.0.0.0", "Host Address: Default 0.0.0.0")
	serverCmd.Flags().StringVarP(&sourcePath, "source-path", "f", "", "Relative Local Path (./data/source) or Google Cloud Storage Bucket Path (gs://XXXXXXX/source/) for Source Files to be Written")
	serverCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written")
	//serverCmd.Flags().StringVarP(&path, "path", "o", "", "Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written")
	//serverCmd.Flags().StringVarP(&storageType, "patstorage-type", "o", "", "LCL for Local or GCS for Google Cloud Storage Bucket")
}
