// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	api "shifter/api"
	"shifter/lib"
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
		log.Println("\033[31m" + `

   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____       ___   ____ ____ ______  ___    ___   ____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/      / _ \ / __// __//_  __/ / _ |  / _ \ /  _/
 ___/ / / / / / __/ /_/  __/ /         / , _// _/ _\ \   / /   / __ | / ___/_/ /
/____/_/ /_/_/_/  \__/\___/_/         /_/|_|/___//___/  /_/   /_/ |_|/_/   /___/

-------------------------------------------------------------------------------------
			` + "\033[0m")

		// Instanciate Shifter Server Instance
		server, err := api.InitServer(serverAddress, serverPort, sourcePath, outputPath)
		if err != nil {
			// Unable to instanciate Shifter HTTP Server
			lib.CLog("error", "Cannot create shifter server: ", err)
		}
		// Start Shifter Server Instance
		err = server.Start()
		if err != nil {
			// Unable to start Shifter HTTP Server
			lib.CLog("error", "Cannot start shifter server: ", err)
		}
	},
}

func init() {
	// TODO - Revisit the Flags and Required Flags, Fix descriptions add options, Add valdations.
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "8082", "Server Port: Default 8082")
	serverCmd.Flags().StringVarP(&serverAddress, "host-address", "a", "0.0.0.0", "Host Address: Default 0.0.0.0")
	serverCmd.Flags().StringVarP(&sourcePath, "source-path", "f", "", "Relative Local Path (./data/source) or Google Cloud Storage Bucket Path (gs://XXXXXXX/source/) for Source Files to be Written")
	serverCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written")
	//serverCmd.Flags().StringVarP(&path, "path", "o", "", "Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written")
	//serverCmd.Flags().StringVarP(&storageType, "patstorage-type", "o", "", "LCL for Local or GCS for Google Cloud Storage Bucket")
}
