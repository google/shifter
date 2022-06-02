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
	"fmt"
	"net/http"
	ops "shifter/ops"

	"github.com/gin-gonic/gin"
)

// Get All Downloadable Objects
func (server *Server) Downloads(ctx *gin.Context) {
	fmt.Println("... Download[s]")
	//downloads := Downloads{}

	// TODO --> Get Dir/Bucket Listing of Objects
	// TODO --> Process Dir/Bucket Listing of Objects to Download Structs

	// using BindJson method to serialize body with struct
	//if err := ctx.BindJSON(&downloads); err != nil {
	//	ctx.AbortWithError(http.StatusBadRequest, err)
	//	return
	//}

	// Construct API Endpoint Response
	r := ResponseDownloads{
		Items:   []*ops.SUID{},
		Message: "Downloads...",
	}
	ctx.JSON(http.StatusOK, r)
}

// Get A Specific Downloadable Object
func (server *Server) Download(ctx *gin.Context) {
	// Validate Project Name has been Provided
	downloadId := ctx.Param("downloadId")
	if downloadId == "" {
		// UUID param required & not found.
		err := errors.New("Download ID Not supplied")
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
	}

	suid, err := ops.ResolveSUID(downloadId)
	if err != nil {
		err := errors.New("Unable to resolve or find Download ID")
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
	}

	r := ResponseDownload{
		SUID:    suid,
		Message: "File Available for Download",
	}
	ctx.JSON(http.StatusOK, r)
}
