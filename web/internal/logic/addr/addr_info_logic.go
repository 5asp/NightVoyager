package addr

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddrInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddrInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddrInfoLogic {
	return &AddrInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddrInfoLogic) AddrInfo(req *types.AddrInfoReq) (resp *types.AddrInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
