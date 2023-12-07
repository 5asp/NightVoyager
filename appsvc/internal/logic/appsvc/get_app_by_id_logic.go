package appsvclogic

import (
	"context"
	"errors"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/model"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"

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
	var app model.App
	if in != nil && in.AppId > 0 {
		query := rel.Where(where.Eq("id", in.AppId))
		if err := l.svcCtx.DB.Find(l.ctx, &app, query); err != nil {
			if !errors.Is(err, rel.ErrNotFound) {
				return nil, err
			}
			return &appsvc.GetAppByIdResp{}, nil
		}
		return &appsvc.GetAppByIdResp{
			Data: &appsvc.App{
				AppId:     int64(app.ID),
				Name:      app.Name,
				Secret:    app.Secret,
				Remark:    app.Remark,
				CreatedAt: app.CreatedAt.Unix(),
				UpdatedAt: app.UpdatedAt.Unix(),
				Status:    int32(app.Status),
				AccountId: int64(app.AccountID),
			},
		}, nil
	}
	return &appsvc.GetAppByIdResp{}, nil
}
