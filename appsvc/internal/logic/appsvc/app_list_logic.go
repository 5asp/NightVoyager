package appsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/model"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"

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
	if in != nil {
		var list []model.App
		var (
			query = rel.Select().SortAsc("id")
		)

		if in.Id > 0 {
			query = query.Where(where.Eq("id", in.Id))
		}
		if in.Name != "" {
			query = query.Where(rel.Like("name", "%"+in.Name+"%"))
		}
		if in.Status > 0 {
			query = query.Where(where.Eq("status", in.Status))
		}
		if in.AccountId > 0 {
			query = query.Where(where.Eq("account_id", in.AccountId))
		}

		if in.CreatedGt > 0 && in.CreatedLt > 0 && in.CreatedLt < in.CreatedGt {
			query = query.Where(rel.Lt("created_at", in.CreatedLt))
			query = query.Where(rel.Gt("created_at", in.CreatedGt))
		}
		query = query.Limit(int(in.PageSize)).Offset(int(in.Page))
		count, err := l.svcCtx.DB.FindAndCountAll(l.ctx, &list, query)
		if err != nil {
			return nil, err
		}
		if count > 0 {
			var apps []*appsvc.App
			for _, v := range list {
				apps = append(apps, &appsvc.App{
					AppId:     int64(v.ID),
					Name:      v.Name,
					Remark:    v.Remark,
					Secret:    v.Secret,
					Status:    int32(v.Status),
					AccountId: int64(v.AccountID),
					CreatedAt: v.CreatedAt.Unix(),
					UpdatedAt: v.UpdatedAt.Unix(),
				})
			}
			return &appsvc.AppListResp{
				List:  apps,
				Total: int64(count),
			}, nil
		}
	}
	return &appsvc.AppListResp{}, nil
}
