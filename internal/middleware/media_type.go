/**
 * Package middleware
 * @file      : media_type.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 14:34
 **/

package middleware

import (
	"fmt"
	"github.com/coverai/api/internal/common/rsp"
	"github.com/coverai/api/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JSONContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		checkUri := false
		for _, targetUri := range AllowedOnlyPostJsonPaths {
			if RegexpUri(targetUri, path) {
				checkUri = true
				break
			}
		}

		method := c.Request.Method
		if checkUri && method != http.MethodPost {
			err := fmt.Errorf("only POST requests are allowed")
			logs.FromContext(c).Error(fmt.Sprintf("path: %s, method: %s, err: %s", path, method, err))
			makeResponseError(c, http.StatusMethodNotAllowed, err)
			return
		}

		contentType := c.ContentType()
		if checkUri && !strings.EqualFold("application/json", contentType) {
			err := fmt.Errorf("Content-Type must be application/json")
			logs.FromContext(c).Error(fmt.Sprintf("path: %s, method: %s, contentType: %s, err: %s", path, method, contentType, err))
			makeResponseError(c, http.StatusUnsupportedMediaType, err)
			return
		}

		c.Next()
	}
}

func makeResponseError(c *gin.Context, httpStatus int, err error) {
	c.AbortWithStatus(httpStatus)
	c.JSON(httpStatus, rsp.MakeResponseEntityError(int32(httpStatus), err))
	c.Abort()
}
