/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	os "shifter/openshift"
)

var (
	namespace string
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
		log.Println("cluster list called")
		log.Println(endpoint)
		log.Println(bearertoken)

		var openshift os.Openshift
		openshift.Endpoint = endpoint
		openshift.AuthToken = bearertoken
		openshift.ListAllResources(namespace)
	},
}

func init() {
	clusterCmd.AddCommand(clusterListCmd)

	clusterListCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Filter to namespace or project")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
