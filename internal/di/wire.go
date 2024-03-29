//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

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
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(string, *conf.Service, *conf.Server, *conf.Data, core.Logger, registry.Registrar, registry.Discovery) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, dao.ProviderSet, newApp))
}
