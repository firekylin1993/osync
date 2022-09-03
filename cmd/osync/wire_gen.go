// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"osync/internal/biz/bizagg"
	"osync/internal/biz/bizedn"
	"osync/internal/conf"
	"osync/internal/data/myagg"
	"osync/internal/data/myedn"
	"osync/internal/data/myotel"
	"osync/internal/db"
	"osync/internal/server"
	"osync/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireDb(data *conf.Data, logger log.Logger) (*db.Data, func(), error) {
	gormDB, err := db.NewMysql(data, logger)
	if err != nil {
		return nil, nil, err
	}
	tidb, err := db.NewTidb(data, logger)
	if err != nil {
		return nil, nil, err
	}
	options := db.NewRedisConf(data)
	client := db.NewRedis(options, logger)
	dbData, err := db.NewData(gormDB, tidb, client)
	if err != nil {
		return nil, nil, err
	}
	return dbData, func() {
	}, nil
}

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *db.Data, logger log.Logger) (*kratos.App, error) {
	greeterService := service.NewGreeterService()
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, nil
}

func wireOtel(contextContext context.Context, confServer *conf.Server, logger log.Logger) (func(), error) {
	client := myotel.NewMetricClient(confServer)
	exporter := myotel.NewMetricExporter(contextContext, client)
	otlptraceClient := myotel.NewTracerClient(confServer)
	otlptraceExporter := myotel.NewTracerExporter(contextContext, otlptraceClient)
	v, err := newOtel(contextContext, confServer, exporter, otlptraceExporter)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func wireOsync(contextContext context.Context, confServer *conf.Server, data *db.Data, logger log.Logger) error {
	ednRepo := myedn.NewEdnRepo(data, confServer, logger)
	ednUsecase := bizedn.NewEdnUsecase(ednRepo, confServer, logger)
	aggRepo := myagg.NewAggRepo(data, logger)
	aggUsecase := bizagg.NewAggUsecase(aggRepo, confServer, logger)
	error2 := newOsync(contextContext, ednUsecase, aggUsecase, logger)
	return error2
}
