package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/appsvc"
	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppLogic {
	return &GetAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppLogic) GetApp(in *appsvc.GetAppReq) (*appsvc.GetAppResp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.GetAppResp{}, nil
}
