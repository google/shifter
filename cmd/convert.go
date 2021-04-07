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
	generators "shifter/generators"
	inputs "shifter/inputs"
)

var (
	inputType string
	input     string
	output    string
	kind      string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert Openshift Resources to Kubernetes native formats",
	Long: `

   _____ __    _ ______           
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /    
/____/_/ /_/_/_/  \__/\___/_/     
                                  

Convert OpenShift resources to kubernetes native formats

Usage: shifter convert -i ./input.yaml -o ./output_dir -k kind
Supply the input file with the -i or --input flag
Supply the output using the -o or --output flag, the directory will be created with the contents of the helm chart.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Shifter - Convert")
		fmt.Println("Converting", inputType, input, "to", kind, output)

		switch inputType {
		case "template":
			t := inputs.Template(input)
			switch kind {
			case "helm":
				generators.Helm(output, t)
			}
		case "yaml":
			t := inputs.Yaml(input)
			switch kind {
			case "yaml":
				generators.Yaml(output, t)
			}
		case "cluster":
			log.Fatal("Openshift resources have not been implemented yet!")

		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inputType, "type", "t", "", "The input type e.g. template, yaml or openshift")
	convertCmd.Flags().StringVarP(&input, "input", "i", "", "Path to the input file to convert, must be in Openshift format")
	convertCmd.Flags().StringVarP(&kind, "kind", "k", "helm", "Output kind options are either: yaml, helm or kpt")
	convertCmd.Flags().StringVarP(&output, "output", "o", "", "Relative path to the output directory for the results on the conversion")
	convertCmd.MarkFlagRequired("input")
	convertCmd.MarkFlagRequired("output")
}
