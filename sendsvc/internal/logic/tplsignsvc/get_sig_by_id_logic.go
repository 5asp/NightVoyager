package tplsignsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSigByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSigByIdLogic {
	return &GetSigByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSigByIdLogic) GetSigById(in *sendsvc.GetSigByIdReq) (*sendsvc.GetSigByIdResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.GetSigByIdResp{}, nil
}
