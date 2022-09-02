// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shifter/lib"
	"shifter/ops"
)

type Download struct {
	Link        string `json:"link"`
	Name        string `json:"name"`
	Uuid        string `json:"uuid"`
	DisplayName string `json:"displayName"`
}

type Downloads struct {
	Items *[]Download `json:"items"`
}

type ResponseDownload struct {
	SUUID   ops.SUUID `json:"suuid"`
	Message string    `json:"message"`
}

type ResponseDownloads struct {
	Items   []*ops.SUUID `json:"items"`
	Message string       `json:"message"`
}

func (server *Server) Download(ctx *gin.Context) {
	downloadId := ctx.Param("downloadId")
	if downloadId == "" {
		err := errors.New("Download ID Not supplied")
		lib.CLog("error", "Downloading", err)
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
	}

	suuid, err := ops.ResolveUUID(downloadId)
	if err != nil {
		err := errors.New("Unable to resolve or find Download ID")
		lib.CLog("error", "Downloading", err)
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
	}

	r := ResponseDownload{
		SUUID:   suuid,
		Message: "File Available for Download",
	}
	ctx.JSON(http.StatusOK, r)
}

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
		Items:   []*ops.SUUID{},
		Message: "Downloads...",
	}
	ctx.JSON(http.StatusOK, r)
}

func (server *Server) DownloadFile(ctx *gin.Context) {
	downloadId := ctx.Param("downloadId")
	if downloadId == "" {
		err := errors.New("Download ID Not supplied")
		lib.CLog("error", "Downloads", err)
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
	}

	suid, err := ops.ResolveUUID(downloadId)
	if err != nil {
		err := errors.New("Unable to resolve or find Download ID")
		lib.CLog("error", "Downloads", err)
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
	}

	fileName := suid.DownloadId + ".zip"
	ctx.File(server.config.serverStorage.outputPath + "/" + fileName)

}
