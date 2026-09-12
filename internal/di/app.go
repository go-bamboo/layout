package di

import (
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/pkg/kratos"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
	"github.com/go-bamboo/pkg/log/sugar"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/rest"
	"github.com/go-bamboo/pkg/rpc"
)

func newApp(srv *conf.Service, logger core.Logger, gs *rpc.Server, hs *rest.Server, r registry.Registrar) *kratos.App {
	app := kratos.New(
		kratos.ID(srv.Id),
		kratos.Name(srv.Name),
		kratos.Version(srv.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.WithOpts(sugar.WithSkip(2))),
		kratos.Server(
			gs,
			hs,
		),
		//kratos.Registrar(r),
	)
	return app
}
