// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user_service.proto

// 包名

package protos

import (
	fmt "fmt"
	math "math"

	proto "google.golang.org/protobuf/proto"

	context "context"

	"github.com/asim/go-micro/v3/api"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	UserReg(ctx context.Context, in *UserModel, opts ...client.CallOption) (*RegResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) UserReg(ctx context.Context, in *UserModel, opts ...client.CallOption) (*RegResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UserReg", in)
	out := new(RegResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	UserReg(context.Context, *UserModel, *RegResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		UserReg(ctx context.Context, in *UserModel, out *RegResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) UserReg(ctx context.Context, in *UserModel, out *RegResponse) error {
	return h.UserServiceHandler.UserReg(ctx, in, out)
}
