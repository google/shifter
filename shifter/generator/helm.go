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

package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"regexp"
	"shifter/lib"

	"gopkg.in/yaml.v3"
	json "k8s.io/apimachinery/pkg/runtime/serializer/json"
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

type Output struct {
	Kind   string
	Output []byte
}

var templates []Output

func mod(o []byte) []byte {
	str1 := string(o)
	var re = regexp.MustCompile(`(?m)\${([^}]*)}`)
	var substitution = "{{.Values.$1}}"
	str1 = re.ReplaceAllString(str1, substitution)
	return []byte(str1)
}

func (generator Generator) Helm(name string, objects []lib.K8sobject, parameters []lib.OSTemplateParams) []lib.Converted {
	var helmChart []lib.Converted

	helmChart = append(helmChart, createChart(name))

	for _, v := range objects {
		//no := strconv.Itoa(k)
		kind := fmt.Sprintf("%v", v.Kind)

		// TODO - Do we need to replace the Log Commented out below?
		//log.Printf("Writing helm template file %x %s", k, kind)

		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := json.NewYAMLSerializer(json.DefaultMetaFactory, nil, nil)

		err := serializer.Encode(v.Object, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		c := mod(buff.Bytes())

		buff.Reset()
		buff.Write(c)

		var resultTemplates lib.Converted
		resultTemplates.Name = kind
		resultTemplates.Path = "/templates/"
		resultTemplates.Payload = *buff
		helmChart = append(helmChart, resultTemplates)
	}

	helmChart = append(helmChart, createValues(parameters))

	return helmChart
}

func createValues(parameters []lib.OSTemplateParams) lib.Converted {
	m := make(map[interface{}]interface{})
	buff := new(bytes.Buffer)

	for _, y := range parameters {
		m[y.Name] = y.Value
	}

	content, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	buff.Write(content)

	var resultValues lib.Converted
	resultValues.Name = "values"
	resultValues.Path = "/"
	resultValues.Payload = *buff

	return resultValues
}

func createChart(name string) lib.Converted {
	var chart Chart
	var version string = "v1.0"

	buff := new(bytes.Buffer)

	chart.ApiVersion = "v2"
	chart.Name = name
	chart.Version = version

	cg, err := yaml.Marshal(&chart)
	if err != nil {
		log.Fatal(err)
	}

	buff.Write(cg)

	var resultChart lib.Converted
	resultChart.Name = "chart"
	resultChart.Path = "/"
	resultChart.Payload = *buff

	return resultChart
}
