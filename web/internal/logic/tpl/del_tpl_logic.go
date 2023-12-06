package tpl

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTplLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTplLogic {
	return &DelTplLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelTplLogic) DelTpl(req *types.TplDelReq) (resp *types.TplDelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
