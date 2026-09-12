package main

import (
	"context"

	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/di"
	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/sugar"
	"github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/uuid"
	"github.com/go-kratos/kratos/v2/encoding/yaml"
	"github.com/spf13/cobra"

	_ "github.com/go-bamboo/pkg/config/file"
	_ "github.com/go-bamboo/pkg/log/file"
	_ "github.com/go-bamboo/pkg/log/std"
	_ "github.com/go-bamboo/pkg/otel/stdout"
	_ "github.com/go-bamboo/pkg/registry/consul"
)

var (
	startCmd = &cobra.Command{
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
	startCmd.Flags().StringVar(&cfg, "conf", "file:///../../configs/conf.yaml", "url for config eg: file:///../../configs/conf.yaml")
}

func run(ctx context.Context) error {
	var bc conf.Bootstrap
	config.Load(cfg, &bc, yaml.Name)

	logger := log.Init(bc.Logger, sugar.WithVersion(Version))
	defer logger.Close()

	// uuid
	id := uuid.New()
	bc.Service.Version = Version
	bc.Service.Id = id

	// consul
	r, d, err := registry.Create(bc.Reg)
	if err != nil {
		log.Errorf("err: %v", err.Error())
		return err
	}

	// trace
	if err := otel.Create(bc.Trace, bc.Service.Name, id, bc.Service.Version); err != nil {
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
