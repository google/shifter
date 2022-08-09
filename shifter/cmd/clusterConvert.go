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
	os "shifter/openshift"
)

// clusterExportCmd represents the clusterExport command
var clusterConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert all resources or resources from a namepace from the cluster.",
	Long: `Convert takes all the resources from a OpenShift cluster endpoint and converts them to the desired output format
on your local disk or GCS bucket.

Examples:
	Convert all resources from a given namespace into yaml files:
	shifter cluster -e $OPENSHIFT_ENDPOINT -t $OPENSHIFT_TOKEN convert -n $NAMESPACE -o yaml ./output/directory/path

	Convert all resources from all namespaces into yaml files:
	shifter cluster -e $OPENSHIFT_ENDPOINT -t $OPENSHIFT_TOKEN convert --all-namespaces -o yaml ./output/directory/path
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
		log.Println("Converting cluster resources")
		procflags := ProcFlags(pFlags)

		var openshift os.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken
		openshift.ConvertNSResources(namespace, procflags, outputPath)
		log.Println("Conversion Complete")
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
