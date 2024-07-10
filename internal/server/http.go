package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/service"
	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metrics"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-bamboo/pkg/rest"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func v1Router(v1 *gin.RouterGroup, s *service.V1Service) {
	chainbot := v1.Group("/chainbot")
	chainbot.POST("webhook", s.ChainBotWebHook)
}

func RegisterApiHTTPServer(httpSrv *http.Server, s *service.V1Service) {
	limiter := bbr.NewLimiter()
	middlewareChain := []middleware.Middleware{
		recovery.Recovery(),
		ratelimit.Server(ratelimit.WithLimiter(limiter)),
		metadata.Server(),
		tracing.Server(),
		metrics.Server(),
		logging.Server(),
	}
	router := gin.Default()
	router.Use(kgin.Middlewares(middlewareChain...))
	api := router.Group("/api")
	api.GET("hello", s.Hello)
	// v1
	v1 := api.Group("/v1")
	v1Router(v1, s)
	httpSrv.HandlePrefix("/", router)
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, admin *service.AdminService, v1Srv *service.V1Service) *rest.Server {
	httpSrv := rest.NewServer(c.Http)
	RegisterApiHTTPServer(httpSrv, v1Srv)
	return httpSrv
}
