package app

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppLsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppLsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLsLogic {
	return &AppLsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppLsLogic) AppLs(req *types.AppLsReq) (resp *types.AppLsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
