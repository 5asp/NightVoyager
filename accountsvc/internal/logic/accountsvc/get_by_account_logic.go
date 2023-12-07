package accountsvclogic

import (
	"context"
	"errors"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/model"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByAccountLogic {
	return &GetByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByAccountLogic) GetByAccount(in *accountsvc.GetByAccountReq) (*accountsvc.GetByAccountResp, error) {
	var account model.Account
	if in != nil && in.Account != "" {
		query := rel.Select("account", "password", "id").Where(where.Eq("account", in.Account))
		if err := l.svcCtx.DB.Find(l.ctx, &account, query); err != nil {
			if !errors.Is(err, rel.ErrNotFound) {
				return nil, err
			}
			return &accountsvc.GetByAccountResp{}, nil
		}
		return &accountsvc.GetByAccountResp{Data: &accountsvc.Account{
			Account:  account.Account,
			Password: account.Password,
			Id:       int64(account.ID),
		}}, nil
	}
	return &accountsvc.GetByAccountResp{}, nil
}
