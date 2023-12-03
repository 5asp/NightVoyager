package logic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/model"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

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
	data := &model.AppInfo{
		Secret:    "11xxxx1",
		Status:    1,
		Remark:    "sas",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	if err := l.svcCtx.DB.Insert(l.ctx, data); err != nil {
		return nil, err
	}
	return &appsvc.CreateAppResp{Id: int64(data.ID)}, nil
}
