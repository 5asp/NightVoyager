package accountlogsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccountLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLogListLogic {
	return &AccountLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccountLogListLogic) AccountLogList(in *accountsvc.AccountLogListReq) (*accountsvc.AccountLogListResp, error) {
	// todo: add your logic here and delete this line

	return &accountsvc.AccountLogListResp{}, nil
}
