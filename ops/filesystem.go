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
	"encoding/json"
	"log"

	//"log"
	"os"
	"path/filepath"
)

func CreateDir(srcPath string) {
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		os.MkdirAll(srcPath, 0700)
	}
}

func GetFilename(path string) string {
	return filepath.Base(path)
}
func GetFileLink(uuid string, path string) string {
	return ("/download/" + uuid + "/" + filepath.Base(path))
}

func GetFiles(uuid string, srcPath string) []File {

	var files []File

	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, File{Filename: GetFilename(path), Link: GetFileLink(uuid, path)})
		return nil
	})
	if err != nil {
		panic(err)
	}

	log.Println(files)
	j, _ := json.Marshal(files)
	log.Println(string(j))
	return files
}

/*

f, err := json.Marshal(file)
if err != nil {
	fmt.Println(err)
	return err
}
fmt.Println(string(f))
files = append(files, f...)
return nil
})
if err != nil {
panic(err)
}

return string(files)



	a, _ := json.Marshal(map[string]int{"foo": 1, "bar": 2, "baz": 3})
fmt.Println(string(a)) // {"bar":2,"baz":3,"foo":1}

	type ResponseYaml2Yaml struct {
		ConvertedFiles []struct {
			FileName string `json:"FileName"`
			Link     string `json:"Link"`
		} `json:"convertedFiles"`
		InputType     string `json:"inputType"`
		Message       string `json:"message"`
		UploadedFiles []struct {
			Filename string `json:"Filename"`
			Header   struct {
				ContentDisposition []string `json:"Content-Disposition"`
				ContentType        []string `json:"Content-Type"`
			} `json:"Header"`
			Size int `json:"Size"`
		} `json:"uploadedFiles"`
		UUID string `json:"uuid"`
	}

	r := ResponseYaml2Yaml{}
	json.Unmarshal([]byte(body), &u)	// Unmarshal


*/

/*var files []string

err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {

convertedFile := ConvertedFile{Filename: GetFilename(path), Link: GetFileLink(uuid, path)}
log.Printf("%+v\n", convertedFile)
byteArray, err := json.Marshal(convertedFile)
if err != nil {
	panic(err)
}
files = append(files, string(byteArray))
return nil
})
if err != nil {
panic(err)
}
return files*/
