package di

import (
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap/zapcore"
)

func newApp(id string, srv *conf.Service, logger zapcore.Core, gs *grpc.Server, hs *http.Server, r registry.Registrar) *kratos.App {
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
