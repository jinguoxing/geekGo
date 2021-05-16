// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/jinguoxing/geekGo/Week04/internal/biz"
	"github.com/jinguoxing/geekGo/Week04/internal/conf"
	"github.com/jinguoxing/geekGo/Week04/internal/data"
	"github.com/jinguoxing/geekGo/Week04/internal/server"
	"github.com/jinguoxing/geekGo/Week04/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, tracerProvider trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUsecase := biz.NewArticleUsecase(articleRepo, logger)
	blogService := service.NewBlogService(articleUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, tracerProvider, blogService)
	grpcServer := server.NewGRPCServer(confServer, tracerProvider, blogService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
