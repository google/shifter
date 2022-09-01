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
	"os"
	"shifter/lib"
	openshift "shifter/openshift"
)

// clusterListCmd represents the clusterList command
var clusterExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports the resources from the source cluster",
	Long: `Export takes the resources 'as-is' from a OpenShift cluster endpoint and exports them in yaml format so the manifests can be fed into the shifter conversion process.'

Examples:
	Export all resources from a given namespace into yaml files:
	shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN export -n $NAMESPACE ./output/directory/path

	Export all resources from all namespaces into yaml files:
	shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN export --all-namespaces ./output/directory/path

`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("\033[31m" + `
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

-----------------------------------
` + "\033[0m")

		if len(args) != 1 {
			lib.CLog("error", "Please specify the destination path")
			os.Exit(1)
		}

		outputPath = args[0]

		lib.CLog("info", "Connecting to cluster: "+endpoint)
		lib.CLog("info", "Converting cluster resources.")

		var openshift openshift.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken

		// Export OpenShift Resources
		err := openshift.ExportNSResources(namespace, outputPath)
		if err != nil {
			// Error: Exporting Resource List
			lib.CLog("error", "Exporting cluster resources: ", err)
			os.Exit(1)
		}

		lib.CLog("info", "Export Complete")
	},
}

func init() {
	clusterCmd.AddCommand(clusterExportCmd)

	clusterExportCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace or Project")
	clusterExportCmd.Flags().BoolVarP(&allnamespaces, "all-namespaces", "", false, "All Namespaces/Projects")
	clusterExportCmd.MarkFlagsMutuallyExclusive("namespace", "all-namespaces")
}
