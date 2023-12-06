package app

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UdAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUdAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UdAppLogic {
	return &UdAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UdAppLogic) UdApp(req *types.AppUdReq) (resp *types.AppUdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
