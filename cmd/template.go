/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Convert openshift templates to helm charts",
	Long: `Convert an openshift template to a helm chart

Usage: shifter template -i ./input.yaml -o ./output_dir
Supply the input file with the -i or --input flag
Supply the output using the -o or --output flag, the directory will be created with the contents of the helm chart.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("templates called")
		fmt.Println(args[0])
		fmt.Println(args[1])
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().BoolP("input", "i", false, "Path to the input file to covert, must be in Openshift format")
	templateCmd.Flags().BoolP("output", "o", false, "Path to the output file for the results on the conversion")
}
