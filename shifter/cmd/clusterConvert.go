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
	"log"
	os "shifter/openshift"
)

// clusterExportCmd represents the clusterExport command
var clusterConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clusterConvert called")
		log.Println("Connecting to cluster: ", endpoint)

		var openshift os.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken
		openshift.ConvertNSResources(namespace)
	},
}

func init() {
	clusterCmd.AddCommand(clusterConvertCmd)

	clusterConvertCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace or Project")
	clusterConvertCmd.Flags().BoolVarP(&allnamespaces, "all-namespaces", "", false, "All Namespaces or Projects")
	clusterConvertCmd.MarkFlagsMutuallyExclusive("namespace", "all-namespaces")
}
