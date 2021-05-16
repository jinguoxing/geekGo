// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"geekGo/Week04/internal/biz"
	"geekGo/Week04/internal/conf"
	"geekGo/Week04/internal/data"
	"geekGo/Week04/internal/server"
	"geekGo/Week04/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
