package appsvclogic

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
	accountID := in.GetAccountId()
	if in != nil && accountID > 0 {
		app := model.App{
			AccountID: int(accountID),
			CreatedAt: time.Now(),
			Status:    1,
			Secret:    utils.CreateUid(),
			Name:      in.GetName(),
		}
		err := l.svcCtx.DB.Insert(l.ctx, &app)
		if err != nil {
			return nil, err
		}
		return &appsvc.CreateAppResp{Id: int64(app.ID)}, nil
	}
	return &appsvc.CreateAppResp{}, nil
}
