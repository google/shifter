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
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// Healthz godoc
// @Summary Shifter Server Healthz.
// @Schemes
// @Description Shifter Server API Endpoint for Operational Health Checks.
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {blob} Blob
// @Router /status/healthz [get]
func Healthz(ctx *gin.Context) {
	// Construct API Endpoint Response
	r := Response_Status_Healthz{}
	r.Timestamp = ""
	r.Status = http.StatusOK
	r.Message = "Shifter Server is reachable."
	// Return JSON API Response
	ctx.JSON(http.StatusOK, r)
}
