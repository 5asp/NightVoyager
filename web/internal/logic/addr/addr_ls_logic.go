package addr

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddrLsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddrLsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddrLsLogic {
	return &AddrLsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddrLsLogic) AddrLs(req *types.AddrLsReq) (resp *types.AddrLsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
