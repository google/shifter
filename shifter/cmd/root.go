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
	"os"
	"shifter/lib"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "shifter",
	Short: "Move your workloads from Openshift to Kubernetes",
	Long: `
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/


Migrate your OpenShift resources to GKE/Anthos`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		lib.CLog("error", "Could not start the shifter cli", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shifter.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			lib.CLog("error", "Unable to find home directory", err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".shifter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		lib.CLog("info", "Using config file: "+viper.ConfigFileUsed())
	}
}

func ProcFlags(input []string) map[string]string {
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
