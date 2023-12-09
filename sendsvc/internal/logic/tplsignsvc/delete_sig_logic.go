package tplsignsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSigLogic {
	return &DeleteSigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSigLogic) DeleteSig(in *sendsvc.DeleteSigReq) (*sendsvc.DeleteSigResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.DeleteSigResp{}, nil
}
