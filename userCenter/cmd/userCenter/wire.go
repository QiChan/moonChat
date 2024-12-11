//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"moonChat/userCenter/internal/biz"
	"moonChat/userCenter/internal/conf"
	"moonChat/userCenter/internal/data"
	"moonChat/userCenter/internal/mq"
	"moonChat/userCenter/internal/server"
	"moonChat/userCenter/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, mq.ProviderSet, newApp))
}
