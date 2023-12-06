package tpl

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TplInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTplInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TplInfoLogic {
	return &TplInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TplInfoLogic) TplInfo(req *types.TplInfoReq) (resp *types.TplInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
