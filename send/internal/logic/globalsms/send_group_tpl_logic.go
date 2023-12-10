package globalsms

import (
	"context"

	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendGroupTplLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendGroupTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendGroupTplLogic {
	return &SendGroupTplLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendGroupTplLogic) SendGroupTpl(req *types.GroupXsendReq) (resp *types.GroupSendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
