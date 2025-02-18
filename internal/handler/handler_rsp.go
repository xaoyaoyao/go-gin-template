/**
 * Package handler
 * @file      : handler_rsp.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 14:20
 **/

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/coverai/api/internal/common/ret"
	"github.com/coverai/api/internal/common/rsp"
	"github.com/coverai/api/internal/common/util"
	"github.com/coverai/api/internal/config"
	"github.com/coverai/api/internal/i18n"
	"github.com/coverai/api/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h HandlerImpl) GetUserApiKey(c *gin.Context) (apiKey *string, apiSecret *string, err error) {
	apiKeyByHeader := c.GetHeader("GO-API-KEY")
	if apiKeyByHeader == "" {
		err := fmt.Errorf("apiKey and apiSecret are illegal")
		logs.FromContext(c).Error(err)
		return nil, nil, err
	}
	delimiter := ":"
	if strings.Contains(apiKeyByHeader, delimiter) {
		parts := strings.Split(apiKeyByHeader, delimiter)
		if len(parts) != 2 {
			err := fmt.Errorf("apiKey and apiSecret are illegal")
			logs.FromContext(c).Error(err)
			return nil, nil, err
		}
		apiKey := parts[0]
		apiSecret := parts[1]
		return &apiKey, &apiSecret, nil
	}
	err = fmt.Errorf("apiKey and apiSecret are illegal")
	logs.FromContext(c).Error(err)
	return nil, nil, err
}

func (h HandlerImpl) IsOpenApiKey(c *gin.Context, apiKey, apiSecret string) bool {
	if apiKey == "" || apiSecret == "" {
		return false
	}
	return apiKey == config.Get().ApiKey && apiSecret == config.Get().SecretKey
}

func (h HandlerImpl) getUserId(c *gin.Context) *string {
	userId, exists := c.Get(util.IdKey)
	if exists {
		uid := fmt.Sprintf("%v", userId)
		if uid == "" {
			return nil
		}
		return &uid
	}
	return nil
}

func (h HandlerImpl) Logger(c *gin.Context, queryParams, pathParams, jsonParams interface{}) {
	userId := h.getUserId(c)
	uid := ""
	if userId != nil {
		uid = *userId
	}
	path := c.Request.URL.Path
	if raw := c.Request.URL.RawQuery; raw != "" {
		path = path + "?" + raw
	}
	method := c.Request.Method
	pathParamsStr, queryParamsStr, jsonParamsStr := "", "", ""
	if pathParams != nil {
		paramsResult, _ := json.Marshal(pathParams)
		if paramsResult != nil {
			pathParamsStr = string(paramsResult)
		}
	}
	if queryParams != nil {
		paramsResult, _ := json.Marshal(queryParams)
		if paramsResult != nil {
			queryParamsStr = string(paramsResult)
		}
	}
	if jsonParams != nil {
		paramsResult, _ := json.Marshal(jsonParams)
		if paramsResult != nil {
			jsonParamsStr = string(paramsResult)
		}
	}
	logs.FromContext(c).Info(fmt.Sprintf("User request parameters. userId: %s, method: %s, path: %s, pathParams: %v, queryParams: %v, jsonParams: %v",
		uid, method, path, pathParamsStr, queryParamsStr, jsonParamsStr))
}

func (h HandlerImpl) makeResponse(c *gin.Context, entity rsp.RetResult) {
	msg := i18n.GetMessage(c, entity.Status)
	if rsp.IsOK(entity) {
		c.JSON(entity.Status, rsp.MakeResponseData(msg, entity.Data))
		return
	}
	if msg != "" {
		entity.Err = fmt.Errorf(msg)
	}
	if entity.Status >= ret.DefaultMinCode {
		c.JSON(http.StatusOK, rsp.MakeResponseEntityError(int32(entity.Status), entity.Err))
		return
	} else if entity.Status == http.StatusInternalServerError {
		c.JSON(http.StatusInternalServerError, rsp.MakeResponseEntityError(int32(entity.Status), entity.Err))
		return
	}
	c.JSON(entity.Status, rsp.MakeResponseEntityError(int32(entity.Status), entity.Err))
}

func (h HandlerImpl) newRetResultOK(c *gin.Context, data interface{}) rsp.RetResult {
	msg := i18n.GetMessage(c, http.StatusOK)
	err := fmt.Errorf(msg)
	return rsp.NewRetResult(http.StatusOK, err, data)
}

func (h HandlerImpl) makeErrorResult(c *gin.Context, code int) rsp.RetResult {
	msg := i18n.GetMessage(c, code)
	err := fmt.Errorf(msg)
	return rsp.NewRetResult(code, err, nil)
}
