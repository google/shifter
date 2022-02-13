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
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// ConvertedFile godoc
// @Summary Download Collection of Converted Files as Zip Archive.
// @Schemes
// @Description Download a collection of Converted Files from the Shifter server by UUID.
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {blob} Blob
// @Router /download/:uuid [get]
func ConvertedFilesArchive(ctx *gin.Context) {

	// Validate URL Params
	// UUID Validation
	uuid := ctx.Param("uuid")
	if uuid == "" {
		// UUID param required & not found.
		err := errors.New("Requested Download URL Path Not Found Error")
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	/*
		TODO
		- Migrate Archive File Path for Download Folder and file to Function
		- Configure File and Folder Path within Server Instantiation Configuration
	*/
	// Construct Archive File Path
	filePath := ("./data/output/" + uuid + "/" + uuid + ".zip")

	// Validate File Path and Archive File Exists
	if _, err := os.Stat(filePath); err == nil {
		// File Exists, Send Download File Response
		ctx.File(filePath)
	} else if errors.Is(err, os.ErrNotExist) {
		// Archive File Does Not Exists, Error
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		log.Fatal("Requested Archive File Not Found Error:", err)
		return
	} else {
		// Archive File Status Unknown, Error
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		log.Fatal("Requested Archive File Status Unknown Error:", err)
		return
	}
}
