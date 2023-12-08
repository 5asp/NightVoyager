package tplsignsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

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

func (l *DeleteSigLogic) DeleteSig(in *tplsvc.DeleteSigReq) (*tplsvc.DeleteSigResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.DeleteSigResp{}, nil
}
