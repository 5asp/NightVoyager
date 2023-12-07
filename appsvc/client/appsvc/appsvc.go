// Code generated by goctl. DO NOT EDIT.
// Source: appsvc.proto

package appsvc

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	App            = appsvc.App
	AppListReq     = appsvc.AppListReq
	AppListResp    = appsvc.AppListResp
	CreateAppReq   = appsvc.CreateAppReq
	CreateAppResp  = appsvc.CreateAppResp
	GetAppByIdReq  = appsvc.GetAppByIdReq
	GetAppByIdResp = appsvc.GetAppByIdResp
	Req            = appsvc.Req
	Resp           = appsvc.Resp

	Appsvc interface {
		GetAppById(ctx context.Context, in *GetAppByIdReq, opts ...grpc.CallOption) (*GetAppByIdResp, error)
		CreateApp(ctx context.Context, in *CreateAppReq, opts ...grpc.CallOption) (*CreateAppResp, error)
		AppList(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*AppListResp, error)
	}

	defaultAppsvc struct {
		cli zrpc.Client
	}
)

func NewAppsvc(cli zrpc.Client) Appsvc {
	return &defaultAppsvc{
		cli: cli,
	}
}

func (m *defaultAppsvc) GetAppById(ctx context.Context, in *GetAppByIdReq, opts ...grpc.CallOption) (*GetAppByIdResp, error) {
	client := appsvc.NewAppsvcClient(m.cli.Conn())
	return client.GetAppById(ctx, in, opts...)
}

func (m *defaultAppsvc) CreateApp(ctx context.Context, in *CreateAppReq, opts ...grpc.CallOption) (*CreateAppResp, error) {
	client := appsvc.NewAppsvcClient(m.cli.Conn())
	return client.CreateApp(ctx, in, opts...)
}

func (m *defaultAppsvc) AppList(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*AppListResp, error) {
	client := appsvc.NewAppsvcClient(m.cli.Conn())
	return client.AppList(ctx, in, opts...)
}