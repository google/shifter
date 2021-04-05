package input

import (
	"fmt"
	gyaml "github.com/ghodss/yaml"
	yaml "gopkg.in/yaml.v3"
	"io"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"os"
	"path/filepath"
	"shifter/processor"
)

type Spec struct {
	Kind string `yaml:"kind"`
}

func Yaml(path string) []runtime.Object {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		t := readMultiFilesInDir(path)
		return t

	case mode.IsRegular():
		t := readMultiDocFile(path)
		return t
	}

	return nil
}

func readMultiFilesInDir(filePath string) []runtime.Object {

	objects := make([]runtime.Object, 0)

	fileList := make([]string, 0)
	e := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
		t := readfile(file)
		objects = append(objects, t)
	}
	return objects
}

func readfile(fileName string) runtime.Object {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	d := yaml.NewDecoder(f)
	doc := make(map[interface{}]interface{})
	err = d.Decode(&doc)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Converting ", doc["kind"])

	val, err := yaml.Marshal(doc)
	if err != nil {
		fmt.Println(err)
	}
	j2, err := gyaml.YAMLToJSON(val)
	if err != nil {
		fmt.Println(err)
	}

	t := processor.Processor(j2, doc["kind"])

	return t
}

func readMultiDocFile(fileName string) []runtime.Object {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	d := yaml.NewDecoder(f)

	objects := make([]runtime.Object, 0)

	for {
		doc := make(map[interface{}]interface{})
		err := d.Decode(&doc)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Converting ", doc["kind"])

		val, err := yaml.Marshal(doc)
		if err != nil {
			fmt.Println(err)
		}
		j2, err := gyaml.YAMLToJSON(val)
		if err != nil {
			fmt.Println(err)
		}

		t := processor.Processor(j2, doc["kind"])
		objects = append(objects, t)

		//fmt.Println(t)
	}
	return objects
}
