package appchannellogic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/model"
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
	if in != nil {
		domain := &model.AppChannel{
			ChannelName:   in.GetName(),
			ChannelKey:    in.GetAppId(),
			ChannelSecret: in.GetSecret(),
			ChannelDomain: in.GetGateway(),
			IsDefault:     int(in.GetIsDefault()),
			CreatedAt:     time.Now(),
		}
		err := l.svcCtx.DB.Insert(l.ctx, domain)
		if err != nil {
			return nil, err
		}
		return &appsvc.CreateChannelResp{Id: domain.ID}, nil
	}

	return &appsvc.CreateChannelResp{}, nil
}
