// Code generated by goctl. DO NOT EDIT.
// Source: accountsvc.proto

package accountsvc

import (
	"context"

	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Account            = accountsvc.Account
	AccountLog         = accountsvc.AccountLog
	AccountLogListReq  = accountsvc.AccountLogListReq
	AccountLogListResp = accountsvc.AccountLogListResp
	CreateReq          = accountsvc.CreateReq
	CreateResp         = accountsvc.CreateResp
	DeleteReq          = accountsvc.DeleteReq
	DeleteResp         = accountsvc.DeleteResp
	GetByAccountReq    = accountsvc.GetByAccountReq
	GetByAccountResp   = accountsvc.GetByAccountResp
	GetByIdReq         = accountsvc.GetByIdReq
	GetByIdResp        = accountsvc.GetByIdResp
	InsertLogReq       = accountsvc.InsertLogReq
	InsertLogResp      = accountsvc.InsertLogResp
	UpdateReq          = accountsvc.UpdateReq
	UpdateResp         = accountsvc.UpdateResp

	AccountSvc interface {
		GetByAccount(ctx context.Context, in *GetByAccountReq, opts ...grpc.CallOption) (*GetByAccountResp, error)
		GetById(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*GetByIdResp, error)
		Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
		Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
		Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
	}

	defaultAccountSvc struct {
		cli zrpc.Client
	}
)

func NewAccountSvc(cli zrpc.Client) AccountSvc {
	return &defaultAccountSvc{
		cli: cli,
	}
}

func (m *defaultAccountSvc) GetByAccount(ctx context.Context, in *GetByAccountReq, opts ...grpc.CallOption) (*GetByAccountResp, error) {
	client := accountsvc.NewAccountSvcClient(m.cli.Conn())
	return client.GetByAccount(ctx, in, opts...)
}

func (m *defaultAccountSvc) GetById(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*GetByIdResp, error) {
	client := accountsvc.NewAccountSvcClient(m.cli.Conn())
	return client.GetById(ctx, in, opts...)
}

func (m *defaultAccountSvc) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	client := accountsvc.NewAccountSvcClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultAccountSvc) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	client := accountsvc.NewAccountSvcClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultAccountSvc) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	client := accountsvc.NewAccountSvcClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}