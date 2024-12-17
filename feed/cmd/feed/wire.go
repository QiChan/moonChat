//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"moonChat/feed/internal/biz"
	"moonChat/feed/internal/conf"
	"moonChat/feed/internal/data"
	"moonChat/feed/internal/mq"
	"moonChat/feed/internal/server"
	"moonChat/feed/internal/service"
	v1 "moonChat/mqInterface/api/msgQueue/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, mq.ProviderSet, v1.ProviderSet, newApp))
}
