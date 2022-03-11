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

package generator

import (
	//"bytes"
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"log"
	"os"
	"regexp"
	"shifter/lib"
	//"strconv"
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

func Helm(path string, objects []lib.K8sobject, parameters []lib.OSTemplateParams, name string) {
	createFolderStruct(path)
	genTemplate(objects, path)
	genValues(parameters, path)
	genChart(path, name, "v1.0.0")
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

func mod(o []byte) []byte {
	str1 := string(o)
	var re = regexp.MustCompile(`(?m)\${([^}]*)}`)
	var substitution = "{{.Values.$1}}"
	str1 = re.ReplaceAllString(str1, substitution)
	return []byte(str1)
}

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	//fmt.Println(string(mod(p)))
	//fmt.Println(len(p))
	p = mod(p)
	fmt.Println(string(p))
	return len(p), nil
}

func genTemplate(objects []lib.K8sobject, path string) {
	for _, y := range objects {
		//no := strconv.Itoa(x)
		//kind := fmt.Sprintf("%v", y.Kind)

		f := new(Writer)
		//f, err := os.Create(path + "/templates/" + no + "-" + kind + ".yaml")
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer f.Close()

		w := bufio.NewWriter(f)
		e := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, nil, nil)

		err := e.Encode(y.Object, w)
		if err != nil {
			fmt.Println(err)
		}
		w.Flush()

		/*
			//get the contents from the object
			content, err := yaml.Marshal(y.Object)
			if err != nil {
				log.Fatal(err)
			}

			//convert the iteration into a string to be used in the filename

			if _, err := file.Write(mod(content)); err != nil {
				log.Fatal(err)
			}
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		*/
	}
}

func genChart(path string, name string, version string) {
	var c Chart
	c.ApiVersion = "v2"
	c.Name = name
	c.Version = version

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

func genValues(parameters []lib.OSTemplateParams, path string) {
	m := make(map[interface{}]interface{})

	for _, y := range parameters {
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
