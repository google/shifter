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
	"os"
	"path/filepath"
)

type JSONResponse struct {
	Value1 string `json:"Filename"`
	Value2 string `json:"Link"`
}

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

func GetFiles(uuid string, srcPath string) []string {

	var files []string

	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, `{"Filename":"`+GetFilename(path)+`", "Link":"`+GetFileLink(uuid, path)+`"}`)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
