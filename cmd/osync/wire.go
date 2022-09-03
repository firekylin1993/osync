//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"osync/internal/biz"
	"osync/internal/conf"
	"osync/internal/data"
	"osync/internal/db"
	"osync/internal/server"
	"osync/internal/service"
)

// wireApp init kratos application.
func wireDb(*conf.Data, log.Logger) (*db.Data, func(), error) {
	panic(
		wire.Build(
			db.ProviderSet,
		),
	)
}

// wireApp init kratos application.
func wireApp(*conf.Server, *db.Data, log.Logger) (*kratos.App, error) {
	panic(
		wire.Build(
			server.ProviderSet,
			service.ProviderSet,
			newApp,
		),
	)
}

func wireOtel(context.Context, *conf.Server, log.Logger) (func(), error) {
	panic(
		wire.Build(
			data.ProviderSet,
			newOtel,
		),
	)
}

func wireOsync(context.Context, *conf.Server, *db.Data, log.Logger) error {
	panic(
		wire.Build(
			biz.ProviderSet,
			data.ProviderSet,
			newOsync,
		),
	)
}
