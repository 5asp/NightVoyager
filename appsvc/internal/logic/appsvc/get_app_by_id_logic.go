package appsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppByIdLogic {
	return &GetAppByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppByIdLogic) GetAppById(in *appsvc.GetAppByIdReq) (*appsvc.GetAppByIdResp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.GetAppByIdResp{}, nil
}
