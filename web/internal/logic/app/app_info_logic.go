package app

import (
	"context"
	"fmt"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppInfoLogic {
	return &AppInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppInfoLogic) AppInfo(req *types.AppInfoReq) (resp *types.AppInfoResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println(l.svcCtx.Config.PageSize)
	return
}
