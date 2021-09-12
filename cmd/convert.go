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
	"github.com/google/shifter/generators"
	"github.com/google/shifter/inputs"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	inputType string
	filename  string
	output    string
	generator string
	pFlags    []string
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

Usage: shifter convert -i ./input.yaml -o ./output_dir -k kind -t kind
Supply the input file or directory of files with the -i or --input flag
Supply the output using the -o or --output flag, the directory will be created with the contents of the helm chart.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
   _____ __    _ ______           
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /    
/____/_/ /_/_/_/  \__/\___/_/     
                                 
----------------------------------------
			`)
		log.Println("Converting", inputType, filename, "to", generator, output)

		flags := procFlags(pFlags)
		log.Println("Processor Flags:", flags)
		switch inputType {
		case "template":
			t := inputs.Template(filename)
			switch generator {
			case "helm":
				generators.Helm(output, t)
			}
		case "yaml":
			t := inputs.Yaml(filename, flags)
			switch generator {
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
	convertCmd.Flags().StringVarP(&inputType, "input-format", "i", "yaml", "Input format. One of: yaml|template")
	convertCmd.Flags().StringVarP(&filename, "filename", "f", "", "Path to the file or directory to convert (contents must be in OpenShift format)")
	convertCmd.Flags().StringVarP(&generator, "output-format", "t", "", "Output format. One of: yaml|helm")
	convertCmd.Flags().StringVarP(&output, "output-path", "o", "", "Relative path to the output directory for the results on the conversion")
	convertCmd.Flags().StringSliceVarP(&pFlags, "pflags", "", []string{}, "Flags passed to the processor")
	convertCmd.MarkFlagRequired("filename")
	convertCmd.MarkFlagRequired("output-path")
}

func procFlags(input []string) map[string]string {
	// Process the inputting processor flags into a map
	m := make(map[string]string)

	for _, f := range input {
		flag := strings.Split(f, "=")
		key := string(flag[0])
		value := string(flag[1])
		m[key] = value
	}

	return m
}
