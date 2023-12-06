package addr

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAddrLogic {
	return &DelAddrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAddrLogic) DelAddr(req *types.AddrDelReq) (resp *types.AddrDelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
