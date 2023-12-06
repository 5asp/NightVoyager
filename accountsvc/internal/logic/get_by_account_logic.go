package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByAccountLogic {
	return &GetByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByAccountLogic) GetByAccount(in *accountsvc.GetByAccountReq) (*accountsvc.GetByAccountResp, error) {
	// todo: add your logic here and delete this line

	return &accountsvc.GetByAccountResp{}, nil
}
