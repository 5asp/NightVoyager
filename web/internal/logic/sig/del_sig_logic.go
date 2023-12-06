package sig

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSigLogic {
	return &DelSigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSigLogic) DelSig(req *types.SigDelReq) (resp *types.SigDelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
