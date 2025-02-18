/**
 * Package service
 * @file      : service.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:00
 **/

package service

import (
	"context"
	"github.com/coverai/api/internal/config"
	"github.com/coverai/api/internal/domain/user"
	"github.com/coverai/api/internal/handler"
	coverAIRouter "github.com/coverai/api/internal/router"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"net/http"
)

type Service struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB) Service {
	// user service
	credential := user.NewCredential()

	httpHandler := handler.NewHandlerImpl(credential)
	httpRouter := coverAIRouter.NewHttpRouter(httpHandler, credential)
	return Service{
		router: httpRouter,
		db:     db,
	}
}

func (s Service) Run(ctx context.Context) error {
	// Run migrations
	//if err := migrations.Run(s.db); err != nil {
	//	logs.FromContext(ctx).Panic(err)
	//	return err
	//}

	// run server
	errgrp, ctx := errgroup.WithContext(ctx)

	server := &http.Server{
		Addr:    config.Get().Addr,
		Handler: s.router,
	}

	errgrp.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	errgrp.Go(func() error {
		<-ctx.Done()
		return server.Close()
	})

	return errgrp.Wait()
}
