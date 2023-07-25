package biz

import (
	"github.com/go-bamboo/layout/internal/rpc"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase, rpc.NewCache)
