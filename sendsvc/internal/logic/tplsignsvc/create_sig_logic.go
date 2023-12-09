package tplsignsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

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

func (l *CreateSigLogic) CreateSig(in *sendsvc.CreateSigReq) (*sendsvc.CreateSigResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.CreateSigResp{}, nil
}
