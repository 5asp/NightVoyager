package globalsms

import (
	"context"

	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendGroupLogic {
	return &SendGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendGroupLogic) SendGroup(req *types.GroupSendReq) (resp *types.GroupSendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
