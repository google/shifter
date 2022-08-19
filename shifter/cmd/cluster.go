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

	"github.com/spf13/cobra"
)

var (
	endpoint      string
	bearertoken   string
	namespace     string
	allnamespaces bool
	csvoutput     bool
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Connect to a running OpenShift cluster.",
	Long: `

   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/



Convert OpenShift resources to kubernetes native formats

Usage: shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN <ACTION>
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
	},
}

func init() {
	rootCmd.AddCommand(clusterCmd)
	clusterCmd.PersistentFlags().StringVarP(&endpoint, "cluster-endpoint", "e", "", "OpenShift cluster endpoint")
	clusterCmd.PersistentFlags().StringVarP(&bearertoken, "token", "t", "", "OpenShift cluster authentication token")
	clusterCmd.MarkPersistentFlagRequired("cluster-endpoint")
	clusterCmd.MarkPersistentFlagRequired("token")
}
