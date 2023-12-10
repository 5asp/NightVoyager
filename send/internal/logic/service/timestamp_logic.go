package service

import (
	"context"

	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TimestampLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTimestampLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TimestampLogic {
	return &TimestampLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TimestampLogic) Timestamp(req *types.Req) (resp *types.TimestampResp, err error) {
	// todo: add your logic here and delete this line

	return
}
