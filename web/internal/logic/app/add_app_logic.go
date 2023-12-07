package app

import (
	"context"
	"fmt"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAppLogic {
	return &AddAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAppLogic) AddApp(req *types.AppAddReq) (resp *types.AppAddResp, err error) {
	fmt.Println(l.ctx.Value("account"))
	return
}
