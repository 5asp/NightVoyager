package tpl

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UdTplLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUdTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UdTplLogic {
	return &UdTplLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UdTplLogic) UdTpl(req *types.TplUdReq) (resp *types.TplUdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
