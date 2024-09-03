package di

import (
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/pkg/kratos"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/rest"
	"github.com/go-bamboo/pkg/rpc"
)

func newApp(id string, srv *conf.Service, logger core.Logger, gs *rpc.Server, hs *rest.Server, r registry.Registrar) *kratos.App {
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(srv.Name),
		kratos.Version(srv.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.NewLogger(logger, 2)),
		kratos.Server(
			gs,
			hs,
		),
		//kratos.Registrar(r),
	)
	return app
}
