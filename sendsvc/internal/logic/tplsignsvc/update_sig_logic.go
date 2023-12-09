package tplsignsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSigLogic {
	return &UpdateSigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSigLogic) UpdateSig(in *sendsvc.UpdateSigReq) (*sendsvc.UpdateSigResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.UpdateSigResp{}, nil
}
