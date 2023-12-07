package auth

import (
	"context"

	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"
	"github.com/aheadIV/NightVoyager/pkg/utils"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegLogic) Reg(req *types.RegReq) (resp *types.RegResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.AccountRPC.GetByAccount(l.ctx, &accountsvc.GetByAccountReq{Account: req.Account})
	if err != nil {
		return nil, err
	}
	accountData := res.GetData()
	if accountData != nil {
		return &types.RegResp{Err: "帐号已经注册"}, nil
	}
	password, _ := utils.CreatePassword(req.Password + l.svcCtx.Config.Auth.AccessSecret)

	_, err = l.svcCtx.AccountRPC.Create(l.ctx, &accountsvc.CreateReq{
		Data: &accountsvc.Account{
			Account:  req.Account,
			Password: password,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.RegResp{}, nil
}
