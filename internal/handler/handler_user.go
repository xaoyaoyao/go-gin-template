/**
 * Package handler
 * @file      : handler_user.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 15:12
 **/

package handler

import (
	"fmt"
	"github.com/coverai/api/internal/common/ret"
	"github.com/coverai/api/internal/common/rsp"
	"github.com/coverai/api/internal/common/util"
	"github.com/coverai/api/internal/http"
	"github.com/coverai/api/internal/httperr"
	"github.com/coverai/api/internal/logs"
	"github.com/gin-gonic/gin"
)

// User signup
// (POST /v1/users/signup)
func (h HandlerImpl) Signup(c *gin.Context, params http.SignupParams) {
	h.Logger(c, params, nil, nil)

	req, done := h.checkRegistration(c)
	if done {
		return
	}
	h.signup(c, req, params)
}

func (h HandlerImpl) checkRegistration(c *gin.Context) (http.SignupJSONBody, bool) {
	apiKey, apiSecret, err := h.GetUserApiKey(c)
	if err != nil {
		logs.FromContext(c).Error(err)
		h.makeResponse(c, rsp.NewBadRequest(err))
		return http.SignupJSONBody{}, true
	} else if apiKey == nil || apiSecret == nil {
		err := fmt.Errorf("apiKey and apiSecret are illegal")
		logs.FromContext(c).Error(err)
		h.makeResponse(c, rsp.NewBadRequest(err))
		return http.SignupJSONBody{}, true
	}
	req := http.SignupJSONBody{}
	if err := c.ShouldBind(&req); err != nil {
		httperr.Respond(c, err)
		return http.SignupJSONBody{}, true
	}
	h.Logger(c, nil, nil, req)

	deviceId := req.DeviceId
	if deviceId == "" {
		err := fmt.Errorf("deviceId is null")
		logs.FromContext(c).Error(err)
		h.makeResponse(c, rsp.NewBadRequest(err))
		return http.SignupJSONBody{}, true
	}
	if !h.IsOpenApiKey(c, *apiKey, *apiSecret) {
		err := fmt.Errorf("apiKey and apiSecret are illegal")
		logs.FromContext(c).Error(err)
		h.makeResponse(c, rsp.NewBadRequest(err))
		return http.SignupJSONBody{}, true
	}
	return req, false
}

func (h HandlerImpl) signup(c *gin.Context, req http.SignupJSONBody, params http.SignupParams) {
	userId := h.CreateOrUpdateUser(c, req)
	credentialResponse, err := h.credential.Auth(c, userId)
	if err != nil || credentialResponse == nil {
		logs.FromContext(c).Error(err)
		h.makeErrorResult(c, ret.Unauthorized.Code())
		return
	}
	credentialVO := http.CredentialVO{
		Id:                    credentialResponse.UserId,
		AccessToken:           credentialResponse.AccessToken,
		ExpiresIn:             credentialResponse.ExpiresIn,
		RefreshToken:          credentialResponse.RefreshToken,
		RefreshTokenExpiresIn: credentialResponse.RefreshTokenExpiresIn,
		Scope:                 &credentialResponse.Scope,
		TokenType:             &credentialResponse.TokenType,
	}
	h.makeResponse(c, h.newRetResultOK(c, credentialVO))
}

func (h HandlerImpl) CreateOrUpdateUser(c *gin.Context, req http.SignupJSONBody) string {
	// TODO 判断设备ID是否已存在，不存在则注册，否则直接返回结果
	return util.NewId()
}
