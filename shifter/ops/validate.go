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

package ops

import (
	"log"
	"github.com/instrumenta/kubeval/kubeval"
)

var (
	config = kubeval.NewDefaultConfig()
)

func ValidateYaml(fileObj *FileObject) (results []kubeval.ValidationResult, err error) {
	schemaCache := kubeval.NewSchemaCache()
	config.FileName = fileObj.Path + "." + fileObj.Ext
	config.SchemaLocation = "https://raw.githubusercontent.com/yannh/kubernetes-json-schema/master"

	log.Println("Validating yaml")

	res, err := kubeval.ValidateWithCache(fileObj.Content.Bytes(), schemaCache, config)
	if err != nil {
		return nil, err
	}

	for  _, r := range res {
		if len(r.Errors) > 0 {
			log.Println("ERROR DETECTED IN YAML")
			log.Println(r)
		}
	}
	return res, nil
}
