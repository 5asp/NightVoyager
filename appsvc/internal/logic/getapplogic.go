package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/model"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"

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
	if in.AppId != 0 {
		var appData model.AppInfo
		err := l.svcCtx.DB.Find(l.ctx, &appData, where.Eq("id", in.AppId))
		switch err {
		case rel.ErrNotFound:
			return &appsvc.GetAppResp{}, nil
		case nil:
			return &appsvc.GetAppResp{
				AppId: int64(appData.ID), Secret: appData.Secret, Remark: appData.Remark, CreatedAt: appData.CreatedAt.Unix(), UpdatedAt: appData.UpdatedAt.Unix(), Status: int32(appData.Status),
			}, nil
		default:
			return nil, err
		}
	}
	return &appsvc.GetAppResp{}, nil
}
