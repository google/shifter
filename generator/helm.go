package generator

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Chart struct {
	ApiVersion  string `yaml:"apiVersion"`
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	KubeVersion string `yaml:"kubeVersion"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
}

type kube struct {
	Parameters []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Required    bool   `yaml:"required"`
		Value       string `yaml:"value"`
	}
	Objects []struct {
		ApiVersion string                 `yaml:"apiVersion"`
		Kind       string                 `yaml:"kind"`
		Metadata   map[string]interface{} //metadata has a unkown structure so we use a generic interface
		Spec       map[string]interface{} //specs are dependent on the kind so we use a generic interface
		Data       map[string]interface{} //specs are dependent on the kind so we use a generic interface
	}
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

	cg, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(cg))
}

func genTemplate() {

}

func genValues() {
}

func Generate(path string, input []byte) {
	createFolderStruct(path)

	chartFile, err := os.Create(path + "/Chart.yaml")
	if err != nil {
		fmt.Println(err)
	}

	valuesFile, err := os.Create(path + "/values.yaml")
	if err != nil {
		fmt.Println(err)
	}

	var data kube
	err = yaml.Unmarshal(input, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	fmt.Println(chartFile)
	fmt.Println(valuesFile)
	genChart()
}
