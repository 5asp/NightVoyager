package tpl

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTplLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTplLogic {
	return &AddTplLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTplLogic) AddTpl(req *types.TplAddReq) (resp *types.TplAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
