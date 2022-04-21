package input

import (
	"fmt"
	gyaml "github.com/ghodss/yaml"
	yaml "gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"shifter/lib"
	"shifter/processor"
	"strings"
)

type Spec struct {
	Kind string `yaml:"kind"`
}

func Yaml(path string, flags map[string]string) (name string, object []lib.K8sobject) {
	file, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Reading in file", file.Name())

	switch mode := file.Mode(); {
	case mode.IsDir():
		fileContent := readMultiFilesInDir(path, flags)
		return file.Name(), fileContent

	case mode.IsRegular():
		fileContent := readMultiDocFile(path, flags)
		return file.Name(), fileContent
	}
	return "", nil
}

func readMultiFilesInDir(filePath string, flags map[string]string) []lib.K8sobject {
	objects := make([]lib.K8sobject, 0)

	fileList := make([]string, 0)
	err := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			fileList = append(fileList, path)
		}
		return err
	})

	if err != nil {
		log.Println(err)
	}

	for _, file := range fileList {
		t := readMultiDocFile(file, flags)
		for _, v := range t {
			objects = append(objects, v)
		}

		//objects = append(objects, t)
	}
	return objects
}

func readMultiDocFile(fileName string, flags map[string]string) []lib.K8sobject {
	nestedQuotedStringHack(fileName)
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	d := yaml.NewDecoder(f)

	objects := make([]lib.K8sobject, 0)

	for {
		doc := make(map[interface{}]interface{})

		err := d.Decode(&doc)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				os.Exit(1)
			}
		}

		if err == io.EOF {
			break
		}
		log.Println("Converting", doc["kind"], "from", fileName)

		val, err := yaml.Marshal(doc)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		jsonBody, err := gyaml.YAMLToJSON(val)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		processedDocument := processor.Processor(jsonBody, doc["kind"], flags)
		objects = append(objects, processedDocument)
	}
	return objects
}

func nestedQuotedStringHack(fileName string) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	str1 := string(input)
	lines := strings.Split(str1, "\n")

	for i, line := range lines {
		found := strings.Contains(line, `\"`)
		if found == true {
			if strings.HasSuffix(lines[i], `'`) == false {
				lines[i] = strings.Replace(lines[i], `"`, `'`, 1)
				lines[i] = strings.TrimSuffix(lines[i], `"`)
				lines[i] = lines[i] + `'`
			}
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
