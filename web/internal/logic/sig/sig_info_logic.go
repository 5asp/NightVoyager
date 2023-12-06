package sig

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SigInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSigInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigInfoLogic {
	return &SigInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SigInfoLogic) SigInfo(req *types.SigInfoReq) (resp *types.SigInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
