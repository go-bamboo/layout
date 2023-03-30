// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.21.12
// source: api/layout/v1/api.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLayoutV1ChainBotWebHook = "/api.layout.v1.LayoutV1/ChainBotWebHook"

type LayoutV1HTTPServer interface {
	ChainBotWebHook(context.Context, *ChainBotWebHookRequest) (*ChainBotWebHookReply, error)
}

func RegisterLayoutV1HTTPServer(s *http.Server, srv LayoutV1HTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/layout/webhook", _LayoutV1_ChainBotWebHook0_HTTP_Handler(srv))
}

func _LayoutV1_ChainBotWebHook0_HTTP_Handler(srv LayoutV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChainBotWebHookRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLayoutV1ChainBotWebHook)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChainBotWebHook(ctx, req.(*ChainBotWebHookRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ChainBotWebHookReply)
		return ctx.Result(200, reply)
	}
}

type LayoutV1HTTPClient interface {
	ChainBotWebHook(ctx context.Context, req *ChainBotWebHookRequest, opts ...http.CallOption) (rsp *ChainBotWebHookReply, err error)
}

type LayoutV1HTTPClientImpl struct {
	cc *http.Client
}

func NewLayoutV1HTTPClient(client *http.Client) LayoutV1HTTPClient {
	return &LayoutV1HTTPClientImpl{client}
}

func (c *LayoutV1HTTPClientImpl) ChainBotWebHook(ctx context.Context, in *ChainBotWebHookRequest, opts ...http.CallOption) (*ChainBotWebHookReply, error) {
	var out ChainBotWebHookReply
	pattern := "/api/v1/layout/webhook"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLayoutV1ChainBotWebHook))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}