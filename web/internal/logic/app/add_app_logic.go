package app

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
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
	if len(req.Name) < 0 {
		return &types.AppAddResp{Err: "应用名称不能为空"}, nil
	}
	accountID := l.ctx.Value("account_id")
	f, err := accountID.(json.Number).Int64()
	if err != nil {
		return nil, errors.New("account not exist.")
	}
	newApp := &appsvc.CreateAppReq{
		AccountId: f,
		Name:      req.Name,
	}
	res, err := l.svcCtx.AppRPC.CreateApp(l.ctx, newApp)
	if err != nil {
		return nil, err
	}
	return &types.AppAddResp{AppID: res.Id}, nil
}
