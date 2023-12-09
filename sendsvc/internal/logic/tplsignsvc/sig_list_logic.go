package tplsignsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SigListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigListLogic {
	return &SigListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SigListLogic) SigList(in *sendsvc.SigListReq) (*sendsvc.SigListResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.SigListResp{}, nil
}
