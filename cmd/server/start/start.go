package start

import (
	"context"
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/di"
	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/uuid"
	"github.com/spf13/cobra"

	_ "github.com/go-bamboo/pkg/config/file"
	_ "github.com/go-bamboo/pkg/log/file"
	_ "github.com/go-bamboo/pkg/log/std"
	_ "github.com/go-bamboo/pkg/otel/stdout"
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
	// cfg is the config remote addr
	cfg string
)

func init() {
	Cmd.Flags().StringVar(&cfg, "conf", "file:///../../configs/conf.yaml", "url for config eg: file:///../../configs/conf.yaml")
}

func run(ctx context.Context) error {
	var bc conf.Bootstrap
	config.Load(cfg, &bc)

	logger := log.Init(bc.Logger)
	defer logger.Close()

	// uuid
	id := uuid.New()

	// consul
	r, d, err := registry.Create(bc.Reg)
	if err != nil {
		log.Errorf("err: %v", err.Error())
		return err
	}

	// trace
	if err := otel.Create(bc.Trace, bc.Service.Name, id); err != nil {
		log.Errorf("err: %v", err.Error())
		return err
	}

	app, closeFunc, err := di.InitApp(id, bc.Service, bc.Server, bc.Data, logger, r, d)
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
