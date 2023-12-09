// Code generated by goctl. DO NOT EDIT.
// Source: tplsvc.proto

package server

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/logic/tplsendsvc"
	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"
)

type TplSendSvcServer struct {
	svcCtx *svc.ServiceContext
	tplsvc.UnimplementedTplSendSvcServer
}

func NewTplSendSvcServer(svcCtx *svc.ServiceContext) *TplSendSvcServer {
	return &TplSendSvcServer{
		svcCtx: svcCtx,
	}
}

func (s *TplSendSvcServer) CreateRec(ctx context.Context, in *tplsvc.CreateRecReq) (*tplsvc.CreateRecResp, error) {
	l := tplsendsvclogic.NewCreateRecLogic(ctx, s.svcCtx)
	return l.CreateRec(in)
}

func (s *TplSendSvcServer) GetRecById(ctx context.Context, in *tplsvc.GetRecByIdReq) (*tplsvc.GetRecByIdResp, error) {
	l := tplsendsvclogic.NewGetRecByIdLogic(ctx, s.svcCtx)
	return l.GetRecById(in)
}