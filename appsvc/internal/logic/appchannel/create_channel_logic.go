package appchannellogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChannelLogic {
	return &CreateChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateChannelLogic) CreateChannel(in *appsvc.CreateChannelReq) (*appsvc.CreateChannelResp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.CreateChannelResp{}, nil
}
