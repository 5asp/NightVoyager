package globalsms

import (
	"context"

	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendTplLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendTplLogic {
	return &SendTplLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendTplLogic) SendTpl(req *types.XsendReq) (resp *types.SendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
