package auth

import (
	"context"

	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegLogic) Reg(req *types.RegReq) (resp *types.RegResp, err error) {
	// todo: add your logic here and delete this line

	return
}
