/**
 * Package handler
 * @file      : handler_init.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 14:17
 **/

package handler

import (
	"github.com/coverai/api/internal/http"
	"github.com/gin-gonic/gin"
)

// Program starts to obtain loading data
// (GET /v1/initialize)
func (h HandlerImpl) InitializationData(c *gin.Context, params http.InitializationDataParams) {
	h.Logger(c, params, nil, nil)

	initializationData := http.InitializationDataVO{}
	h.makeResponse(c, h.newRetResultOK(c, initializationData))
}
