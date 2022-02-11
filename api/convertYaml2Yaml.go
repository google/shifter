/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	ops "shifter/ops"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Ep_ConvertYaml2Yaml(c *gin.Context) {

	// Create API Unique RUN ID
	uuid := uuid.New().String()

	// Create Raw Input Folder if not Exists
	srcPath := ("./data/raw/" + uuid + "/")
	ops.CreateDir(srcPath)
	// Create Raw Output Folder if not Exists
	dstPath := ("./data/output/" + uuid + "/")
	ops.CreateDir(dstPath)

	form, _ := c.MultipartForm()
	files := form.File["multiplefiles"]
	for _, file := range files {
		log.Print(file.Filename)
		dst := path.Join(srcPath, file.Filename)
		//Upload files to the specified directory
		c.SaveUploadedFile(file, dst)
	}

	// Run the Conversion Operation
	ops.Convert("yaml", srcPath, "yaml", dstPath, make(map[string]string))
	ops.Archive(dstPath, (dstPath + "/" + uuid + ".zip"))

	// ops.Archive(dstPath)

	// Response
	c.JSON(http.StatusOK, gin.H{
		"inputType":      "yaml",
		"uuid":           string(uuid),
		"convertedFiles": json.Marshal(`[{"Filename":"Tim"},{"Filename":"Gary"}]`),
		"uploadedFiles":  files,
		"message":        "YAML files generated.",
	})
}
