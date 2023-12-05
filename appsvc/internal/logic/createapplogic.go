package logic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/model"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/aheadIV/NightVoyager/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAppLogic {
	return &CreateAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAppLogic) CreateApp(in *appsvc.CreateAppReq) (*appsvc.CreateAppResp, error) {
	// todo: add your logic here and delete this line
	secret := utils.CreateUid()
	data := &model.AppInfo{
		Secret:    secret,
		Status:    1,
		Remark:    "11",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	if err := l.svcCtx.DB.Insert(l.ctx, data); err != nil {
		return nil, err
	}
	return &appsvc.CreateAppResp{Id: int64(data.ID), Secret: secret}, nil
}
