package tpl

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TplLsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTplLsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TplLsLogic {
	return &TplLsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TplLsLogic) TplLs(req *types.TplLsReq) (resp *types.TplLsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
