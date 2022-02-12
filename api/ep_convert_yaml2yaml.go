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
<<<<<<< HEAD
=======
	"encoding/json"
>>>>>>> e0be23d (v0.3 API Endpoints, Basic Error response, GIN Server Structured Correctly, Request Validation and Minor Error Handling, 3 Routes in Play)
	"log"
	"net/http"
	"path"
	ops "shifter/ops"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

<<<<<<< HEAD
// @BasePath /api/v1

// Yaml2Yaml godoc
// @Summary Openshift Manifest to Kubernetes Manifest.
// @Schemes
// @Description Convert Openshift Yaml Manifest files into Kubernetes Manifest files.
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {json} Response_Convert_Yaml2Yaml
// @Router /convert/yaml/yaml [post]
func Yaml2Yaml(ctx *gin.Context) {
=======
<<<<<<< HEAD:api/convertYaml2Yaml.go
func Ep_ConvertYaml2Yaml(c *gin.Context) {
=======
func Yaml2Yaml(ctx *gin.Context) {
>>>>>>> e0be23d (v0.3 API Endpoints, Basic Error response, GIN Server Structured Correctly, Request Validation and Minor Error Handling, 3 Routes in Play):api/ep_convert_yaml2yaml.go
>>>>>>> e0be23d (v0.3 API Endpoints, Basic Error response, GIN Server Structured Correctly, Request Validation and Minor Error Handling, 3 Routes in Play)

	// Create API Unique RUN ID
	uuid := uuid.New().String()

	// Create Raw Input Folder if not Exists
	srcPath := ("./data/raw/" + uuid + "/")
	ops.CreateDir(srcPath)

	// Create Raw Output Folder if not Existsk
	dstPath := ("./data/output/" + uuid + "/")
	ops.CreateDir(dstPath)

	// Validate that Request Contains at least One File
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		log.Fatal("Multipart Form Error:", err)
		return
	}

	// Collect Files from Multipart Form.
	files := form.File["multiplefiles"]
	for _, file := range files {
<<<<<<< HEAD
		log.Println(file.Filename)
=======
>>>>>>> e0be23d (v0.3 API Endpoints, Basic Error response, GIN Server Structured Correctly, Request Validation and Minor Error Handling, 3 Routes in Play)
		//Upload files to the specified directory
		ctx.SaveUploadedFile(file, path.Join(srcPath, file.Filename))
	}

	/*
		TODO
		- Add Errors Handling to ops.Convert(),
		- Catch Conversion Errors
		- Respond to API with Error Message JSON.
	*/
	// Run the Conversion Operation
	ops.Convert("yaml", srcPath, "yaml", dstPath, make(map[string]string))

	/*
		TODO
		- Add Errors Handling to ops.Archive(),
		- Catch Archive Errors,
		- Respond to API with Error Message JSON.
	*/
	// Run the Archive Operation
	ops.Archive(dstPath, (dstPath + "/" + uuid + ".zip"))

<<<<<<< HEAD
=======
<<<<<<< HEAD:api/convertYaml2Yaml.go
	// ops.Archive(dstPath)

	// Response
	c.JSON(http.StatusOK, gin.H{
		"inputType":      "yaml",
		"uuid":           string(uuid),
		"convertedFiles": json.Marshal(`[{"Filename":"Tim"},{"Filename":"Gary"}]`),
		"uploadedFiles":  files,
		"message":        "YAML files generated.",
	})
=======
>>>>>>> e0be23d (v0.3 API Endpoints, Basic Error response, GIN Server Structured Correctly, Request Validation and Minor Error Handling, 3 Routes in Play)
	// Construct API Endpoint Response
	r := Response_Convert_Yaml2Yaml{}
	r.InputType = "yaml"
	r.UUID = string(uuid)
	r.ConvertedFiles = ops.GetFiles(uuid, dstPath)
	r.UploadedFiles = files
	r.Message = "YAML files generated."
	// Return JSON API Response
	ctx.JSON(http.StatusOK, r)
<<<<<<< HEAD
=======
>>>>>>> e0be23d (v0.3 API Endpoints, Basic Error response, GIN Server Structured Correctly, Request Validation and Minor Error Handling, 3 Routes in Play):api/ep_convert_yaml2yaml.go
>>>>>>> e0be23d (v0.3 API Endpoints, Basic Error response, GIN Server Structured Correctly, Request Validation and Minor Error Handling, 3 Routes in Play)
}
