package app

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAppLogic {
	return &DelAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAppLogic) DelApp(req *types.AppDelReq) (resp *types.AppDelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
