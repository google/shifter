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
	"github.com/spf13/cobra"
	"log"
	openshift "shifter/openshift"
)

// clusterListCmd represents the clusterList command
var clusterExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports the resources from the source cluster",
	Long: `Export takes the resources 'as-is' from a OpenShift cluster endpoint and exports them in yaml format so the manifests can be fed into the shifter conversion process.'

Examples:
	Export all resources from a given namespace into yaml files:
	shifter cluster -e $OPENSHIFT_ENDPOINT -t $OPENSHIFT_TOKEN export -n $NAMESPACE ./output/directory/path

	Export all resources from all namespaces into yaml files:
	shifter cluster -e $OPENSHIFT_ENDPOINT -t $OPENSHIFT_TOKEN export --all-namespaces ./output/directory/path

`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(`
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

----------------------------------------
`)
		if len(args) != 1 {
			log.Fatal("Please specify the destination path.")
		}

		outputPath = args[0]

		log.Println("Connecting to cluster: ", endpoint)
		log.Println("Exporting cluster resources")
		var openshift openshift.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken
		openshift.ExportNSResources(namespace, outputPath)
		log.Println("Export Complete")
	},
}

func init() {
	clusterCmd.AddCommand(clusterExportCmd)

	clusterExportCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace or Project")
	clusterExportCmd.Flags().BoolVarP(&allnamespaces, "all-namespaces", "", false, "All Namespaces/Projects")
	clusterExportCmd.MarkFlagsMutuallyExclusive("namespace", "all-namespaces")
}
