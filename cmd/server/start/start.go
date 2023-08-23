package start

import (
	"context"
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/di"
	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/registry"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "start",
		Short: "start",
		Long:  `entry`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context())
		},
	}
)

func run(ctx context.Context) error {
	var bc conf.Bootstrap
	config.Load(&bc)

	logger := log.Init(bc.Logger)
	defer logger.Close()

	// uuid
	id, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("err: %v", err)
		return err
	}

	// consul
	r, d := registry.New(bc.Reg)

	// trace
	tp, err := otel.NewProvider(bc.Trace, bc.Service.Name, id.String())
	if err != nil {
		log.Errorf("err: %v", err.Error())
		return err
	}
	defer tp.Close()

	app, closeFunc, err := di.InitApp(id.String(), bc.Service, bc.Server, bc.Data, logger, r, d)
	if err != nil {
		log.Errorf("err: %v", err.Error())
		return err
	}
	defer closeFunc()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Errorf("err: %v", err.Error())
		return err
	}
	return nil
}
