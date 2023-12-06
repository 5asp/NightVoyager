package addr

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddrLogic {
	return &AddAddrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddrLogic) AddAddr(req *types.AddrAddReq) (resp *types.AddrAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
