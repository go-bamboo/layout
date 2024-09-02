package server

import (
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/service"
	"github.com/go-bamboo/pkg/rest"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, admin *service.AdminService, v1Srv *service.V1Service) *rest.Server {
	httpSrv := rest.NewServer(c.Http)
	RegisterApiHTTPServer(httpSrv, v1Srv)
	return httpSrv
}
