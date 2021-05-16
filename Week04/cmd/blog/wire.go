// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinguoxing/geekGo/Week04/internal/biz"
	"github.com/jinguoxing/geekGo/Week04/internal/conf"
	"github.com/jinguoxing/geekGo/Week04/internal/data"
	"github.com/jinguoxing/geekGo/Week04/internal/server"
	"github.com/jinguoxing/geekGo/Week04/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
