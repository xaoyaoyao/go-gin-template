// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Program starts to obtain loading data
	// (GET /v1/initialize)
	InitializationData(c *gin.Context, params InitializationDataParams)
	// User signup
	// (POST /v1/users/signup)
	Signup(c *gin.Context, params SignupParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// InitializationData operation middleware
func (siw *ServerInterfaceWrapper) InitializationData(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params InitializationDataParams

	// ------------- Optional query parameter "queryParams" -------------

	err = runtime.BindQueryParameter("form", true, false, "queryParams", c.Request.URL.Query(), &params.QueryParams)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter queryParams: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.InitializationData(c, params)
}

// Signup operation middleware
func (siw *ServerInterfaceWrapper) Signup(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params SignupParams

	headers := c.Request.Header

	// ------------- Required header parameter "GO-API-KEY" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("GO-API-KEY")]; found {
		var GOAPIKEY string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for GO-API-KEY, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "GO-API-KEY", valueList[0], &GOAPIKEY, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: true})
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter GO-API-KEY: %w", err), http.StatusBadRequest)
			return
		}

		params.GOAPIKEY = GOAPIKEY

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Header parameter GO-API-KEY is required, but not found"), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Signup(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/v1/initialize", wrapper.InitializationData)
	router.POST(options.BaseURL+"/v1/users/signup", wrapper.Signup)
}
