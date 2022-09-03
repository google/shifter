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
	"shifter/openshift"
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
		log.Println("\033[31m" + `
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

----------------------------------------
` + "\033[0m")

		if len(args) != 1 {
			lib.CLog("error", "Please specify the destination path")
			os.Exit(1)
		}

		outputPath = args[0]

		lib.CLog("info", "Connecting to cluster: "+endpoint)
		lib.CLog("info", "Converting cluster resources.")
		procflags := ProcFlags(pFlags)

		var openshift openshift.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken

		// Convert OpenShift Resources
		err := openshift.ConvertNSResources(namespace, procflags, outputPath)
		if err != nil {
			lib.CLog("error", "Converting cluster resources", err)
			os.Exit(1)
		}
		lib.CLog("info", "Conversion Complete")
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
