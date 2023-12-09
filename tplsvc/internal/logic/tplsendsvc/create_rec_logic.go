package tplsendsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRecLogic {
	return &CreateRecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRecLogic) CreateRec(in *tplsvc.CreateRecReq) (*tplsvc.CreateRecResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.CreateRecResp{}, nil
}