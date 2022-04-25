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
	"bytes"
	"io"
	"net/http"
	"path/filepath"

	ops "shifter/ops"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
func (server *Server) Yaml2Yaml(ctx *gin.Context) {

	// Create API Unique RUN ID
	uuid := uuid.New().String()

	// Validate that Request Contains at least One File
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create a New Instance of Converter for File Upload
	converter := &ops.Converter{}
	converter.UUID = uuid

	// Collect Files from Multipart Form.
	files := form.File["multiplefiles"]
	for _, file := range files {

		// Read File Contents
		fileContents, _ := file.Open()
		byteContainer, err := io.ReadAll(fileContents)
		if err != nil {
			// If Unable to Read File into Byte Array
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		//var sourceFiles []*ops.FileObject
		// Create FileObject for Each uploaded File.
		fileObj := &ops.FileObject{
			StorageType:   server.config.serverStorage.storageType,
			SourcePath:    (server.config.serverStorage.sourcePath + "/" + uuid + "/" + file.Filename),
			Ext:           filepath.Ext(file.Filename),
			Content:       *bytes.NewBuffer(byteContainer),
			ContentLength: len(byteContainer),
		}

		// Add File Object to Array of Files
		converter.SourceFiles = append(converter.SourceFiles, fileObj)
		converter.WriteSourceFiles()
	}

	// Create a New Instance of the Converter to Convert the Files
	converter = ops.NewConverter(ops.YAML, (server.config.serverStorage.sourcePath + "/" + uuid), ops.YAML, (server.config.serverStorage.outputPath + "/" + uuid), make(map[string]string))
	converter.UUID = uuid

	// Run the Conversions
	converter.LoadSourceFiles()
	converter.ConvertFiles()

	// Construct API Endpoint Response
	r := Response_Convert_Yaml2Yaml{}
	r.InputType = ops.YAML
	r.UUID = string(uuid)
	r.ConvertedFiles = converter.BuildDownloadFiles()
	r.UploadedFiles = files
	r.Message = "YAML files generated."
	// Return JSON API Response
	ctx.JSON(http.StatusOK, r)
}
