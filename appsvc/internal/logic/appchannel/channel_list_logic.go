package appchannellogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChannelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelListLogic {
	return &ChannelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChannelListLogic) ChannelList(in *appsvc.ChannelListReq) (*appsvc.ChannelListResp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.ChannelListResp{}, nil
}
