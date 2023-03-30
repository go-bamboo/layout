// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.21.12
// source: api/layout/admin/admin.proto

package admin

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

const OperationQuantBotAdminCreateBot = "/api.layout.admin.QuantBotAdmin/CreateBot"
const OperationQuantBotAdminDeleteBot = "/api.layout.admin.QuantBotAdmin/DeleteBot"
const OperationQuantBotAdminGetBot = "/api.layout.admin.QuantBotAdmin/GetBot"
const OperationQuantBotAdminListBot = "/api.layout.admin.QuantBotAdmin/ListBot"
const OperationQuantBotAdminPageBot = "/api.layout.admin.QuantBotAdmin/PageBot"
const OperationQuantBotAdminUpdateBot = "/api.layout.admin.QuantBotAdmin/UpdateBot"

type QuantBotAdminHTTPServer interface {
	CreateBot(context.Context, *CreateBotRequest) (*CreateBotReply, error)
	DeleteBot(context.Context, *DeleteBotRequest) (*DeleteBotReply, error)
	GetBot(context.Context, *GetBotRequest) (*GetBotReply, error)
	ListBot(context.Context, *ListBotRequest) (*ListBotReply, error)
	PageBot(context.Context, *PageBotRequest) (*PageBotReply, error)
	UpdateBot(context.Context, *UpdateBotRequest) (*UpdateBotReply, error)
}

func RegisterQuantBotAdminHTTPServer(s *http.Server, srv QuantBotAdminHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/layout/botPage", _QuantBotAdmin_PageBot0_HTTP_Handler(srv))
	r.GET("/admin/v1/layout/botList", _QuantBotAdmin_ListBot0_HTTP_Handler(srv))
	r.GET("/admin/v1/layout/bot/{id}", _QuantBotAdmin_GetBot0_HTTP_Handler(srv))
	r.POST("/admin/v1/layout/bot", _QuantBotAdmin_CreateBot0_HTTP_Handler(srv))
	r.PUT("/admin/v1/layout/bot", _QuantBotAdmin_UpdateBot0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/layout/bot/{ids}", _QuantBotAdmin_DeleteBot0_HTTP_Handler(srv))
}

func _QuantBotAdmin_PageBot0_HTTP_Handler(srv QuantBotAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageBotRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQuantBotAdminPageBot)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PageBot(ctx, req.(*PageBotRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageBotReply)
		return ctx.Result(200, reply)
	}
}

func _QuantBotAdmin_ListBot0_HTTP_Handler(srv QuantBotAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListBotRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQuantBotAdminListBot)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBot(ctx, req.(*ListBotRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListBotReply)
		return ctx.Result(200, reply)
	}
}

func _QuantBotAdmin_GetBot0_HTTP_Handler(srv QuantBotAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBotRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQuantBotAdminGetBot)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBot(ctx, req.(*GetBotRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBotReply)
		return ctx.Result(200, reply)
	}
}

func _QuantBotAdmin_CreateBot0_HTTP_Handler(srv QuantBotAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBotRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQuantBotAdminCreateBot)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBot(ctx, req.(*CreateBotRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBotReply)
		return ctx.Result(200, reply)
	}
}

func _QuantBotAdmin_UpdateBot0_HTTP_Handler(srv QuantBotAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBotRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQuantBotAdminUpdateBot)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBot(ctx, req.(*UpdateBotRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateBotReply)
		return ctx.Result(200, reply)
	}
}

func _QuantBotAdmin_DeleteBot0_HTTP_Handler(srv QuantBotAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteBotRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQuantBotAdminDeleteBot)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBot(ctx, req.(*DeleteBotRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteBotReply)
		return ctx.Result(200, reply)
	}
}

type QuantBotAdminHTTPClient interface {
	CreateBot(ctx context.Context, req *CreateBotRequest, opts ...http.CallOption) (rsp *CreateBotReply, err error)
	DeleteBot(ctx context.Context, req *DeleteBotRequest, opts ...http.CallOption) (rsp *DeleteBotReply, err error)
	GetBot(ctx context.Context, req *GetBotRequest, opts ...http.CallOption) (rsp *GetBotReply, err error)
	ListBot(ctx context.Context, req *ListBotRequest, opts ...http.CallOption) (rsp *ListBotReply, err error)
	PageBot(ctx context.Context, req *PageBotRequest, opts ...http.CallOption) (rsp *PageBotReply, err error)
	UpdateBot(ctx context.Context, req *UpdateBotRequest, opts ...http.CallOption) (rsp *UpdateBotReply, err error)
}

type QuantBotAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewQuantBotAdminHTTPClient(client *http.Client) QuantBotAdminHTTPClient {
	return &QuantBotAdminHTTPClientImpl{client}
}

func (c *QuantBotAdminHTTPClientImpl) CreateBot(ctx context.Context, in *CreateBotRequest, opts ...http.CallOption) (*CreateBotReply, error) {
	var out CreateBotReply
	pattern := "/admin/v1/layout/bot"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQuantBotAdminCreateBot))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuantBotAdminHTTPClientImpl) DeleteBot(ctx context.Context, in *DeleteBotRequest, opts ...http.CallOption) (*DeleteBotReply, error) {
	var out DeleteBotReply
	pattern := "/admin/v1/layout/bot/{ids}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQuantBotAdminDeleteBot))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuantBotAdminHTTPClientImpl) GetBot(ctx context.Context, in *GetBotRequest, opts ...http.CallOption) (*GetBotReply, error) {
	var out GetBotReply
	pattern := "/admin/v1/layout/bot/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQuantBotAdminGetBot))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuantBotAdminHTTPClientImpl) ListBot(ctx context.Context, in *ListBotRequest, opts ...http.CallOption) (*ListBotReply, error) {
	var out ListBotReply
	pattern := "/admin/v1/layout/botList"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQuantBotAdminListBot))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuantBotAdminHTTPClientImpl) PageBot(ctx context.Context, in *PageBotRequest, opts ...http.CallOption) (*PageBotReply, error) {
	var out PageBotReply
	pattern := "/admin/v1/layout/botPage"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQuantBotAdminPageBot))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuantBotAdminHTTPClientImpl) UpdateBot(ctx context.Context, in *UpdateBotRequest, opts ...http.CallOption) (*UpdateBotReply, error) {
	var out UpdateBotReply
	pattern := "/admin/v1/layout/bot"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQuantBotAdminUpdateBot))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}