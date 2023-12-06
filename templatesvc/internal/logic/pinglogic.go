package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/templatesvc/internal/svc"
	"github.com/aheadIV/NightVoyager/templatesvc/types/templatesvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *templatesvc.Request) (*templatesvc.Response, error) {
	// todo: add your logic here and delete this line

	return &templatesvc.Response{}, nil
}
