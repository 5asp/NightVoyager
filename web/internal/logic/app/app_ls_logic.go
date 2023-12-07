package app

import (
	"context"
	"fmt"

	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppLsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppLsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLsLogic {
	return &AppLsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppLsLogic) AppLs(req *types.AppLsReq) (resp *types.AppLsResp, err error) {
	res, err := l.svcCtx.AppRPC.AppList(l.ctx, &appsvc.AppListReq{
		PageSize: l.svcCtx.Config.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var list []types.App
	if res != nil && res.Total > 0 {
		fmt.Println("++++++++++++++++++++++++++++++++++=")
		for _, v := range res.List {
			list = append(list, types.App{
				Id:        int(v.AppId),
				Name:      v.Name,
				Remark:    v.Remark,
				Secret:    v.Secret,
				Status:    int(v.Status),
				AccountId: int(v.AccountId),
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}
		fmt.Println(list)

		return &types.AppLsResp{Total: int(res.GetTotal()), List: list}, nil
	}
	return &types.AppLsResp{}, nil
}
