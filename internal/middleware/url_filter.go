/**
 * Package middleware
 * @file      : url_filter.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 14:35
 **/

package middleware

import (
	"fmt"
	"github.com/coverai/api/internal/common/ret"
	"github.com/coverai/api/internal/common/rsp"
	"github.com/coverai/api/internal/common/util"
	"github.com/coverai/api/internal/domain/user"
	"github.com/coverai/api/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func URLFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpRequest := c.Request
		path := httpRequest.URL.Path
		method := httpRequest.Method
		logs.FromContext(c).Infof("path: %s, method: %s", path, method)

		if IsDetailURL(path, method, c) || IsAllowedAllUriByParameters(path) {
			c.Next()
			return
		}

		if !allowedPathsMap[path] {
			c.AbortWithStatus(http.StatusNotFound)
			err := fmt.Errorf("not found")
			c.JSON(http.StatusNotFound, rsp.MakeResponseEntityNotFound(err))
			return
		}

		c.Next()
	}
}

func IsDetailURL(path string, method string, c *gin.Context) bool {
	if method == "GET" {
		for key, value := range DetailPathMapByGet {
			if RegexpUri(key, path) && c.Param(value) != "" {
				return true
			}
		}
	} else if method == "HEAD" || method == "PATCH" || method == "OPTIONS" {
		for key, value := range DetailPathMapByHeadOrPatch {
			if RegexpUri(key, path) && c.Param(value) != "" {
				return true
			}
		}
	} else if method == "POST" {
		for key, value := range DetailPathMapByPost {
			if RegexpUri(key, path) && c.Param(value) != "" {
				return true
			}
		}
	}
	return false
}

func Auth(credential *user.Credential) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpRequest := c.Request
		path := httpRequest.URL.Path
		if IsIgnoreAuthorizationUri(path) {
			c.Next()
			return
		}
		token := util.ExtractToken(httpRequest.Header.Get(util.AuthorizationKey))
		isAllow := isValidTokenAndSetContext(c, token, path, credential)
		if isAllow {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			rsp.MakeResponseEntity(int32(ret.Forbidden.Code()), "Unauthorized."),
		)
	}
}

func isValidTokenAndSetContext(c *gin.Context, token string, path string, credential *user.Credential) (isAllow bool) {
	isIgnorePath := IsIgnoreAuthorizationUri(path)
	if isIgnorePath {
		return true
	}
	if token == "" {
		return false
	}
	claims, err := credential.CheckToken(c, token)
	if err != nil || claims == nil {
		if err == nil {
			err = fmt.Errorf("claims is nil")
		}
		logs.FromContext(c).Error(err)
		return false
	}
	return true
}

func IsIgnoreAuthorizationUri(path string) bool {
	for _, targetUri := range IgnoreAuthorizationUriPaths {
		if RegexpUri(targetUri, path) {
			return true
		}
	}
	return false
}

func IsAllowedAllUriByParameters(path string) bool {
	for _, targetUri := range AllowedAllPathsByParameters {
		if RegexpUri(targetUri, path) {
			return true
		}
	}
	return false
}

func GetPath(version string, path string) string {
	return util.RootPath + version + path
}

func RegexpUri(sourceUri, targetUri string) bool {
	regex := regexp.MustCompile(util.RootPath + `/v\d+` + sourceUri)
	return regex.MatchString(targetUri)
}
