package generator

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"reflect"
	"strconv"
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
	Objects []object
}

type object struct {
	ApiVersion string                 `yaml:"apiVersion"`
	Kind       string                 `yaml:"kind"`
	Metadata   map[string]interface{} //metadata has a unkown structure so we use a generic interface
	Spec       map[string]interface{} `yaml:"spec,omitempty"` //specs are dependent on the kind so we use a generic interface
	Data       map[string]interface{} `yaml:"data,omitempty"` //specs are dependent on the kind so we use a generic interface
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

func walk(v interface{}) {
	fmt.Println(v)

	val := reflect.ValueOf(v)
	fmt.Println(val)
	fmt.Println(val.Kind())

	//	if val.Kind() == reflect.Map {
	//		for _, e := range val.MapKeys() {
	//			fmt.Println(e)
	//			walk(e)
	//		}
	//	}

	switch val.Kind() {
	case reflect.Map:
		for _, e := range val.MapKeys() {
			fmt.Println(e)
			walk(e)
		}
	default:
		fmt.Println(val)
	}
}

func mod(o object) {
	for k, v := range o.Spec {
		fmt.Println(k, v)
		walk(v)
		//matched, err := regexp.Match(`\${([^}]+)}`, []byte(v))
		//fmt.Println(matched, err)
		fmt.Println("\n")
	}
}

func genChart(path string) {
	var c Chart
	c.ApiVersion = "v2"
	c.Name = "test"
	c.Version = "1.0.0"

	cg, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	chartFile, err := os.Create(path + "/Chart.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := chartFile.Write(cg); err != nil {
		log.Fatal(err)
	}
	if err := chartFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func genTemplate(objects kube, path string) {
	for x, y := range objects.Objects {
		mod(y)
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

func genValues(parameters kube, path string) {
	m := make(map[interface{}]interface{})

	for _, y := range parameters.Parameters {
		m[y.Name] = y.Value
	}

	content, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(path + "/values.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.Write(content); err != nil {
		log.Fatal(err)
	}
}

func Generate(path string, input []byte) {
	createFolderStruct(path)

	var data kube
	err := yaml.Unmarshal(input, &data)
	if err != nil {
		fmt.Println(err)
	}

	genTemplate(data, path)
	genValues(data, path)
	genChart(path)
}
