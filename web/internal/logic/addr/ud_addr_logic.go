package addr

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UdAddrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUdAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UdAddrLogic {
	return &UdAddrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UdAddrLogic) UdAddr(req *types.AddrUdReq) (resp *types.AddrUdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
