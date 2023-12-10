package appchannellogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChannelByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelByIdLogic {
	return &GetChannelByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChannelByIdLogic) GetChannelById(in *appsvc.GetChannelByIdReq) (*appsvc.GetChannelByIdResp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.GetChannelByIdResp{}, nil
}
