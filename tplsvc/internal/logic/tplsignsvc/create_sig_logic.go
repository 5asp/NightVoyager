package tplsignsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSigLogic {
	return &CreateSigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSigLogic) CreateSig(in *tplsvc.CreateSigReq) (*tplsvc.CreateSigResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.CreateSigResp{}, nil
}
