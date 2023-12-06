package sig

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSigLogic {
	return &AddSigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSigLogic) AddSig(req *types.SigAddReq) (resp *types.SigAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
