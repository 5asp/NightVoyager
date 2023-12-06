package sig

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SigLsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSigLsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigLsLogic {
	return &SigLsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SigLsLogic) SigLs(req *types.SigLsReq) (resp *types.SigLsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
