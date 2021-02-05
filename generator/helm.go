package generator

import (
		"gopkg.in/yaml.v2"
	"fmt"
	"os"
)

type Chart struct {
	ApiVersion  string `yaml:"apiVersion"`
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	KubeVersion string `yaml:"kubeVersion"`
	Description string `yaml:"description"`
	Type string `yaml:"type"`
}

// This needs to be broken out into a module!
func createFolderStruct(path string) {
	var chartsFldr string
	var templatesFldr string

	chartsFldr = path + "/charts"
	templatesFldr = path + "/templates"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}
	if _, err := os.Stat(chartsFldr); os.IsNotExist(err) {
		os.Mkdir(chartsFldr, 0700)
	}
	if _, err := os.Stat(templatesFldr); os.IsNotExist(err) {
		os.Mkdir(templatesFldr, 0700)
	}
}

func genChart() {
 var c Chart
	c.ApiVersion = "v2"
	c.Name = "test"
        c.Version = "1.0.0"
	
	cg, err  := yaml.Marshal(&c)
	if err != nil {
	fmt.Println(err)
	}

fmt.Println(string(cg))
}

func CreateChart(path string) {
	createFolderStruct(path)

	chartFile, err := os.Create(path + "/Chart.yaml")
	if err != nil {
		fmt.Println(err)
	}

	valuesFile, err := os.Create(path + "/Values.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(chartFile)
	fmt.Println(valuesFile)
	genChart()
}
