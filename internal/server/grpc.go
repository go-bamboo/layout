package server

import (
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/pkg/rpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server) *rpc.Server {
	srv := rpc.NewServer(c.Grpc)
	return srv
}
