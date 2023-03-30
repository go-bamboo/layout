package server

import (
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/pkg/rpc"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.uber.org/zap/zapcore"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger zapcore.Core) *grpc.Server {
	srv := rpc.NewServer(c.Grpc)
	return srv
}
