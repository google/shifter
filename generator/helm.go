package generator

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
"strconv"
"log"
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
		Spec       map[string]interface{} `yaml:"spec,omitempty"` //specs are dependent on the kind so we use a generic interface
		Data       map[string]interface{} `yaml:"data,omitempty"` //specs are dependent on the kind so we use a generic interface
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

func genChart(path string) {
	var c Chart
	c.ApiVersion = "v2"
	c.Name = "test"
	c.Version = "1.0.0"

	cg, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	chartFile, err := os.Create(path + "/Chart.yaml")
	if err != nil {
		fmt.Println(err)
	}

	if _, err := chartFile.Write(cg); err != nil {
		fmt.Println(err)
	}
	if err := chartFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func genTemplate(objects kube, path string) {
	for x, y := range objects.Objects {
		content, err := yaml.Marshal(y)
		if err != nil {
			log.Fatal(err)
		}
		no := strconv.Itoa(x)
		file, err := os.Create(path + "/templates/" + no + "-" + y.Kind + ".yaml")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := file.Write(content); err != nil {
			log.Fatal(err)
		}
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func genValues() {
}

func Generate(path string, input []byte) {
	createFolderStruct(path)

	valuesFile, err := os.Create(path + "/values.yaml")
	if err != nil {
		fmt.Println(err)
	}

	var data kube
	err = yaml.Unmarshal(input, &data)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range data.Objects {
		fmt.Println(k, v)
		fmt.Println("-------")
	}



	genTemplate(data, path)

	fmt.Println(valuesFile)
	genChart(path)
}
