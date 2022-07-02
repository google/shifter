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

var (
	namespace     string
	allnamespaces bool
	csvoutput     bool
)

// clusterListCmd represents the clusterList command
var clusterListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(`
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

----------------------------------------
`)
		log.Println("Connecting to cluster: ", endpoint)

		var openshift os.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken
		openshift.ListNSResources(csvoutput, namespace)
	},
}

func init() {
	clusterCmd.AddCommand(clusterListCmd)

	clusterListCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace or Project")
	clusterListCmd.Flags().BoolVarP(&allnamespaces, "all-namespaces", "", false, "All Namespaces/Projects")
	clusterListCmd.Flags().BoolVarP(&csvoutput, "csv", "", false, "CSV Output")
	clusterListCmd.MarkFlagsMutuallyExclusive("namespace", "all-namespaces")
}
