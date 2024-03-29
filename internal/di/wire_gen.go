// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/go-bamboo/layout/internal/biz"
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/dao"
	"github.com/go-bamboo/layout/internal/server"
	"github.com/go-bamboo/layout/internal/service"
	"github.com/go-bamboo/pkg/log/core"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-kratos/kratos/v2"
	registry2 "github.com/go-kratos/kratos/v2/registry"
)

// Injectors from wire.go:

// initApp init kratos application.
func InitApp(string2 string, confService *conf.Service, confServer *conf.Server, data *conf.Data, logger core.Logger, registrar registry.Registrar, discovery registry2.Discovery) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(confServer)
	query := dao.NewQuery(data)
	daoDao, cleanup, err := dao.New(data, query)
	if err != nil {
		return nil, nil, err
	}
	greeterUsecase, err := biz.NewGreeterUsecase(query, daoDao)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	adminService := service.NewAdminService(greeterUsecase)
	v1Service := service.NewV1Service(greeterUsecase)
	httpServer := server.NewHTTPServer(confServer, adminService, v1Service)
	app := newApp(string2, confService, logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
