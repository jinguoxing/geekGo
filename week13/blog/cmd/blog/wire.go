// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"geekGo/week13/blog/internal/biz"
	"geekGo/week13/blog/internal/conf"
	"geekGo/week13/blog/internal/data"
	"geekGo/week13/blog/internal/server"
	"geekGo/week13/blog/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
