package sendlogsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogListLogic {
	return &SendLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendLogListLogic) SendLogList(in *appsvc.Req) (*appsvc.Resp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.Resp{}, nil
}
