package generator

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

type Chart struct {
	ApiVersion  string `yaml:"apiVersion"`
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	KubeVersion string `yaml:"kubeVersion"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Icon        string `yaml:"icon"`
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

func walk(input interface{}) {

	va := reflect.ValueOf(&input).Elem()
	val := va.Elem()

	switch val.Kind() {
	case reflect.Map:
		fmt.Println("***** MAP *******")
		fmt.Println(val)

		for i, k := range val.MapKeys() {
			v := val.MapIndex(k)
			fmt.Println(i, k, v)
			fmt.Println("!!!!!!!! CAN SET: ", v.CanSet())
			walk(v.Interface())

		}
		fmt.Println("CAN SET: ", val.CanSet())
	case reflect.String:
		fmt.Println("***** STRING *******")
		fmt.Println(val)
		fmt.Println("CAN SET: ", val.CanSet())
	default:
		fmt.Println("**** OTHER: ", val.Kind())
		fmt.Println(val)
		fmt.Println("CAN SET: ", val.CanSet())

	}

	/*

		val := reflect.ValueOf(input)

		if val.Kind() == reflect.Map {
			//fmt.Println("MAP!!!!")
			for i, k := range val.MapKeys() {
				v := val.MapIndex(k)
				fmt.Println(i, k, v)
				fmt.Println(val.CanSet())
				walk(v.Interface())

			}
		} else {

			switch val.Kind() {

			case reflect.String:
				fmt.Println("*** FOUND STRING ***")
				fmt.Println(val)
				fmt.Println(val.CanSet())
				fmt.Println("********************")

			default:
				fmt.Println(val.Kind())
			}

		}
	*/
}

func mod(o []byte) []byte {
	str1 := string(o)

	var re = regexp.MustCompile(`(?m)\${([^}]*)}`)

	var substitution = "{{.Values.$1}}"
	str1 = re.ReplaceAllString(str1, substitution)
	//str1 = strings.Replace(str1, "${", "{{.Values.", -1)
	//str1 = strings.Replace(str1, "}", "}}", -1)
	return []byte(str1)
}

func genTemplate(objects kube, path string) {
	for x, y := range objects.Objects {

		//get the contents from the object
		content, err := yaml.Marshal(y)
		if err != nil {
			log.Fatal(err)
		}

		//convert the iteration into a string to be used in the filename
		no := strconv.Itoa(x)
		file, err := os.Create(path + "/templates/" + no + "-" + y.Kind + ".yaml")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := file.Write(mod(content)); err != nil {
			log.Fatal(err)
		}
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
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
