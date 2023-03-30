package main

import (
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/di"
	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/otel"
	"github.com/google/uuid"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// Branch is current branch name the code is built off.
	Branch string
	// Revision is the short commit hash of source tree.
	Revision string
	// BuildDate is the date when the binary was built.
	BuildDate string
)

func main() {
	var bc conf.Bootstrap
	config.Load(&bc)

	logger := log.Init(bc.Logger)
	defer logger.Close()

	// uuid
	id, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}

	// consul
	//r := registry.New(consulClient)

	// trace
	tp, err := otel.NewProvider(bc.Trace, bc.Service.Name, id.String())
	if err != nil {
		log.Errorf("err: %v", err.Error())
		return
	}
	defer tp.Close()

	app, closeFunc, err := di.InitApp(id.String(), bc.Service, bc.Server, bc.Data, bc.Market, logger, nil, nil)
	if err != nil {
		log.Errorf("err: %v", err.Error())
		return
	}
	defer closeFunc()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Errorf("err: %v", err.Error())
		return
	}
}
