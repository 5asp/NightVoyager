package sig

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UdSigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUdSigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UdSigLogic {
	return &UdSigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UdSigLogic) UdSig(req *types.SigUdReq) (resp *types.SigUdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
