/**
 * Package router
 * @file      : router.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:07
 **/

package router

import (
	"context"
	"github.com/coverai/api/internal/common/util"
	"github.com/coverai/api/internal/domain/user"
	"github.com/coverai/api/internal/handler"
	"github.com/coverai/api/internal/http"
	"github.com/coverai/api/internal/i18n"
	"github.com/coverai/api/internal/logs"
	"github.com/coverai/api/internal/middleware"
	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"time"
)

func NewHttpRouter(
	handler handler.HandlerImpl,
	credential *user.Credential,
) *gin.Engine {
	// gin new engine
	router := gin.New()
	// i18n
	router.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		DefaultLanguage:  language.English,
		FormatBundleFile: "yaml",
		AcceptLanguage:   []language.Tag{language.English, language.Chinese},
		RootPath:         "./localize/",
		UnmarshalFunc:    yaml.Unmarshal,
		Loader: &ginI18n.EmbedLoader{
			FS: i18n.FS,
		},
	}), ginI18n.WithGetLngHandle(
		// Obtain the language from the request parameters and set it as the default language.
		func(context *gin.Context, defaultLng string) string {
			lng := context.Query("language")
			if lng == "" {
				return defaultLng
			}
			return lng
		},
	)))
	router.Use(gin.Recovery())
	// set cors
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Authorization", "X-Requested-With", "X-Request-ID", "X-HTTP-Method-Override", "Upload-Length", "Upload-Offset", "Tus-Resumable", "Upload-Metadata",
			"Upload-Defer-Length", "Upload-Concat", "User-Agent", "Referrer", "Origin", "Content-Type", "Content-Length", "X-SESSION-ID",
		},
		ExposeHeaders: []string{
			"Content-Disposition", "Upload-Offset", "Location", "Upload-Length", "Tus-Version", "Tus-Resumable", "Tus-Max-Size", "Tus-Extension", "Upload-Metadata",
			"Upload-Defer-Length", "Upload-Concat", "Location", "Upload-Offset", "Upload-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))
	// set health
	router.GET("/health", HealthHandler)
	router.GET("/healthcheck", HealthCheck)
	router.POST("/health/:code", UpdateHealthHandler)
	// block URL
	router.Use(middleware.URLFilter())
	router.Use(middleware.JSONContentTypeMiddleware())
	// register api
	apiRouter := router.Group(util.RootPath)
	apiHandler := apiRouter.Group("")
	// add auth && logger
	apiHandler.Use(middleware.Auth(credential))
	apiHandler.Use(middleware.Logger())
	// Registration API
	registerHandlers(handler, apiHandler)
	// logging
	logs.FromContext(context.Background()).Info("service started")
	return router
}

func registerHandlers(handler handler.HandlerImpl, apiHandler *gin.RouterGroup) {
	http.RegisterHandlers(apiHandler, handler)
}
