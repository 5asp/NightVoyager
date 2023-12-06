package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sigsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sigsvc/sigsvc"

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

func (l *PingLogic) Ping(in *sigsvc.Request) (*sigsvc.Response, error) {
	// todo: add your logic here and delete this line

	return &sigsvc.Response{}, nil
}
