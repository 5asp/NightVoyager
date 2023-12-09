// Code generated by goctl. DO NOT EDIT.
// Source: tplsvc.proto

package tplsignsvc

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRecReq   = tplsvc.CreateRecReq
	CreateRecResp  = tplsvc.CreateRecResp
	CreateSigReq   = tplsvc.CreateSigReq
	CreateSigResp  = tplsvc.CreateSigResp
	CreateTplReq   = tplsvc.CreateTplReq
	CreateTplResp  = tplsvc.CreateTplResp
	DeletTplResp   = tplsvc.DeletTplResp
	DeleteSigReq   = tplsvc.DeleteSigReq
	DeleteSigResp  = tplsvc.DeleteSigResp
	DeleteTplReq   = tplsvc.DeleteTplReq
	GetRecByIdReq  = tplsvc.GetRecByIdReq
	GetRecByIdResp = tplsvc.GetRecByIdResp
	GetSigByIdReq  = tplsvc.GetSigByIdReq
	GetSigByIdResp = tplsvc.GetSigByIdResp
	GetTplByIdReq  = tplsvc.GetTplByIdReq
	GetTplByIdResp = tplsvc.GetTplByIdResp
	Rec            = tplsvc.Rec
	SigListReq     = tplsvc.SigListReq
	SigListResp    = tplsvc.SigListResp
	Sign           = tplsvc.Sign
	Tpl            = tplsvc.Tpl
	TplListReq     = tplsvc.TplListReq
	TplListResp    = tplsvc.TplListResp
	UpdateSigReq   = tplsvc.UpdateSigReq
	UpdateSigResp  = tplsvc.UpdateSigResp
	UpdateTplReq   = tplsvc.UpdateTplReq
	UpdateTplResp  = tplsvc.UpdateTplResp

	TplSignSvc interface {
		GetSigById(ctx context.Context, in *GetSigByIdReq, opts ...grpc.CallOption) (*GetSigByIdResp, error)
		CreateSig(ctx context.Context, in *CreateSigReq, opts ...grpc.CallOption) (*CreateSigResp, error)
		UpdateSig(ctx context.Context, in *UpdateSigReq, opts ...grpc.CallOption) (*UpdateSigResp, error)
		DeleteSig(ctx context.Context, in *DeleteSigReq, opts ...grpc.CallOption) (*DeleteSigResp, error)
		SigList(ctx context.Context, in *SigListReq, opts ...grpc.CallOption) (*SigListResp, error)
	}

	defaultTplSignSvc struct {
		cli zrpc.Client
	}
)

func NewTplSignSvc(cli zrpc.Client) TplSignSvc {
	return &defaultTplSignSvc{
		cli: cli,
	}
}

func (m *defaultTplSignSvc) GetSigById(ctx context.Context, in *GetSigByIdReq, opts ...grpc.CallOption) (*GetSigByIdResp, error) {
	client := tplsvc.NewTplSignSvcClient(m.cli.Conn())
	return client.GetSigById(ctx, in, opts...)
}

func (m *defaultTplSignSvc) CreateSig(ctx context.Context, in *CreateSigReq, opts ...grpc.CallOption) (*CreateSigResp, error) {
	client := tplsvc.NewTplSignSvcClient(m.cli.Conn())
	return client.CreateSig(ctx, in, opts...)
}

func (m *defaultTplSignSvc) UpdateSig(ctx context.Context, in *UpdateSigReq, opts ...grpc.CallOption) (*UpdateSigResp, error) {
	client := tplsvc.NewTplSignSvcClient(m.cli.Conn())
	return client.UpdateSig(ctx, in, opts...)
}

func (m *defaultTplSignSvc) DeleteSig(ctx context.Context, in *DeleteSigReq, opts ...grpc.CallOption) (*DeleteSigResp, error) {
	client := tplsvc.NewTplSignSvcClient(m.cli.Conn())
	return client.DeleteSig(ctx, in, opts...)
}

func (m *defaultTplSignSvc) SigList(ctx context.Context, in *SigListReq, opts ...grpc.CallOption) (*SigListResp, error) {
	client := tplsvc.NewTplSignSvcClient(m.cli.Conn())
	return client.SigList(ctx, in, opts...)
}