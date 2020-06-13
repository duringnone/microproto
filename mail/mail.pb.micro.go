// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: mail/mail.proto

package mail

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Mail service

func NewMailEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Mail service

type MailService interface {
	SendMail(ctx context.Context, in *SendMailRequest, opts ...client.CallOption) (*SendMailResponse, error)
	// 邮件模板
	GetMailConfigList(ctx context.Context, in *GetMailConfigListRequest, opts ...client.CallOption) (*GetMailConfigListResponse, error)
	AddMailConfig(ctx context.Context, in *AddMailConfigRequest, opts ...client.CallOption) (*AddMailConfigResponse, error)
	UpdateMailConfig(ctx context.Context, in *UpdateMailConfigRequest, opts ...client.CallOption) (*UpdateMailConfigResponse, error)
}

type mailService struct {
	c    client.Client
	name string
}

func NewMailService(name string, c client.Client) MailService {
	return &mailService{
		c:    c,
		name: name,
	}
}

func (c *mailService) SendMail(ctx context.Context, in *SendMailRequest, opts ...client.CallOption) (*SendMailResponse, error) {
	req := c.c.NewRequest(c.name, "Mail.SendMail", in)
	out := new(SendMailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailService) GetMailConfigList(ctx context.Context, in *GetMailConfigListRequest, opts ...client.CallOption) (*GetMailConfigListResponse, error) {
	req := c.c.NewRequest(c.name, "Mail.GetMailConfigList", in)
	out := new(GetMailConfigListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailService) AddMailConfig(ctx context.Context, in *AddMailConfigRequest, opts ...client.CallOption) (*AddMailConfigResponse, error) {
	req := c.c.NewRequest(c.name, "Mail.AddMailConfig", in)
	out := new(AddMailConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailService) UpdateMailConfig(ctx context.Context, in *UpdateMailConfigRequest, opts ...client.CallOption) (*UpdateMailConfigResponse, error) {
	req := c.c.NewRequest(c.name, "Mail.UpdateMailConfig", in)
	out := new(UpdateMailConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mail service

type MailHandler interface {
	SendMail(context.Context, *SendMailRequest, *SendMailResponse) error
	// 邮件模板
	GetMailConfigList(context.Context, *GetMailConfigListRequest, *GetMailConfigListResponse) error
	AddMailConfig(context.Context, *AddMailConfigRequest, *AddMailConfigResponse) error
	UpdateMailConfig(context.Context, *UpdateMailConfigRequest, *UpdateMailConfigResponse) error
}

func RegisterMailHandler(s server.Server, hdlr MailHandler, opts ...server.HandlerOption) error {
	type mail interface {
		SendMail(ctx context.Context, in *SendMailRequest, out *SendMailResponse) error
		GetMailConfigList(ctx context.Context, in *GetMailConfigListRequest, out *GetMailConfigListResponse) error
		AddMailConfig(ctx context.Context, in *AddMailConfigRequest, out *AddMailConfigResponse) error
		UpdateMailConfig(ctx context.Context, in *UpdateMailConfigRequest, out *UpdateMailConfigResponse) error
	}
	type Mail struct {
		mail
	}
	h := &mailHandler{hdlr}
	return s.Handle(s.NewHandler(&Mail{h}, opts...))
}

type mailHandler struct {
	MailHandler
}

func (h *mailHandler) SendMail(ctx context.Context, in *SendMailRequest, out *SendMailResponse) error {
	return h.MailHandler.SendMail(ctx, in, out)
}

func (h *mailHandler) GetMailConfigList(ctx context.Context, in *GetMailConfigListRequest, out *GetMailConfigListResponse) error {
	return h.MailHandler.GetMailConfigList(ctx, in, out)
}

func (h *mailHandler) AddMailConfig(ctx context.Context, in *AddMailConfigRequest, out *AddMailConfigResponse) error {
	return h.MailHandler.AddMailConfig(ctx, in, out)
}

func (h *mailHandler) UpdateMailConfig(ctx context.Context, in *UpdateMailConfigRequest, out *UpdateMailConfigResponse) error {
	return h.MailHandler.UpdateMailConfig(ctx, in, out)
}
