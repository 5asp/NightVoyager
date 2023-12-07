package appsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppListLogic {
	return &AppListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppListLogic) AppList(in *appsvc.AppListReq) (*appsvc.AppListResp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.AppListResp{}, nil
}
