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
	"log"
	os "shifter/openshift"

	"github.com/spf13/cobra"
)

// clusterExportCmd represents the clusterExport command
var clusterConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert all OpenShift resources from a namepace from the OpenShift cluster.",
	Long: `Convert takes all the resources from a OpenShift cluster endpoint and converts them to the desired output format on your local disk or GCS bucket.

Examples:
	Convert all resources from a given namespace into yaml files:
	shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN convert -n $NAMESPACE -o yaml ./output/directory/path

	Convert all resources from all namespaces into yaml files:
	shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN convert --all-namespaces -o yaml ./output/directory/path
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
			log.Fatal("🧰 ❌ ERROR: Please specify the destination path.")
		}

		outputPath = args[0]

		log.Printf("🧰 💡 INFO: Connecting to cluster: '%s'", endpoint)
		log.Printf("🧰 💡 INFO: Converting cluster resources.")
		procflags := ProcFlags(pFlags)

		var openshift os.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken
		// Convert OpenShift Resources
		err := openshift.ConvertNSResources(namespace, procflags, outputPath)
		if err != nil {
			// Error: Converting Resource List
			log.Fatal(fmt.Sprintf("🧰 ❌ ERROR: Converting Resource List: '%s'. ", err.Error()))
		}
		log.Println("Conversion Complete")
		log.Printf("🧰 ✅ SUCCESS: Conversion Complete")
		log.Printf("👋 INFO: Thats all Folks.. Bye Bye!")
	},
}

func init() {
	clusterCmd.AddCommand(clusterConvertCmd)

	clusterConvertCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace or Project")
	clusterConvertCmd.Flags().BoolVarP(&allnamespaces, "all-namespaces", "", false, "All Namespaces or Projects")
	clusterConvertCmd.Flags().StringP("output-format", "o", "yaml", "Output format")

	clusterConvertCmd.MarkFlagRequired("output-format")
	clusterConvertCmd.MarkFlagsMutuallyExclusive("namespace", "all-namespaces")

	clusterConvertCmd.Flags().StringSliceVarP(&pFlags, "pflags", "p", []string{}, "Flags passed to the processor")
}
