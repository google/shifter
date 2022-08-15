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
	ops "shifter/ops"

	"github.com/spf13/cobra"
)

var (
	inputType  string
	sourcePath string
	outputPath string
	useIstio   bool
	generator  string
	pFlags     []string
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

Usage: shifter convert -i yaml -k yaml source/folder/or/file output/folder/or/file
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

		// Welcome Banners
		log.Printf("👋 INFO: Welcome to Shifter Converter")
		log.Printf("🎬 INFO: Let's Start Shifting...")

		if len(args) != 2 {
			log.Fatal("🧰 ❌ ERROR: Please specify source and destination arguments.")
		}
		sourcePath = args[0]
		outputPath = args[1]

		log.Println("🧰 🚀 Converting", inputType, sourcePath, "to", generator, outputPath)
		procflags := ProcFlags(pFlags)
		if useIstio == true {
			procflags["istio"] = "true"
		}

		// Create new Shifter Converter
		con, err := ops.NewConverter(inputType, sourcePath, generator, outputPath, procflags)
		if err != nil {
			// Error: Creating New Shifter Converter
			log.Printf("🧰 ❌ ERROR: Creating Shifter Converter, Unable to continue.")
			log.Printf("🧰 ❌ ERROR: '%s'.", err)
			return
		} else {
			// Success: Creating New Shifter Converter
			log.Printf("🧰 ✅ SUCCESS: Creating Shifter Converter")
		}
		err = con.ConvertFiles()
		if err != nil {
			// Error: Converting Files
			log.Printf("🧰 ❌ ERROR: Converterting provided files, Unable to continue.")
			log.Printf("🧰 ❌ ERROR: '%s'.", err)
			return
		} else {
			// Success: Files Converted
			log.Printf("🧰 ✅ SUCCESS: Provided files converted")
		}
		log.Printf("✅ SUCCESS: Shifter Conversion Complete")
		log.Printf("👋 INFO: Thats all Folks.. Bye Bye!")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inputType, "input-format", "i", "yaml", "Input format. One of: yaml|template")
	convertCmd.Flags().StringVarP(&generator, "output-format", "o", "", "Output format. One of: yaml|helm")
	convertCmd.Flags().StringSliceVarP(&pFlags, "pflags", "p", []string{}, "Flags passed to the processor")
	convertCmd.Flags().BoolVarP(&useIstio, "istio", "m", false, "Use istio for routes conversion")
	convertCmd.MarkFlagRequired("input-format")
	convertCmd.MarkFlagRequired("output-format")
}
