package appchannellogic

import (
	"context"
	"errors"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/model"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"

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
	if in != nil && in.AppId > 0 {
		var appChannel model.AppChannel
		query := rel.Where(where.Eq("id", in.AppId))
		if err := l.svcCtx.DB.Find(l.ctx, &appChannel, query); err != nil {
			if !errors.Is(err, rel.ErrNotFound) {
				return nil, err
			}
			return &appsvc.GetChannelByIdResp{}, nil
		}
		return &appsvc.GetChannelByIdResp{
			Data: &appsvc.Channel{
				Id:            appChannel.ID,
				ChannelName:   appChannel.ChannelName,
				ChannelKey:    appChannel.ChannelKey,
				ChannelDomain: appChannel.ChannelDomain,
				ChannelSecret: appChannel.ChannelSecret,
				IsDefault:     int64(appChannel.IsDefault),
				Quota:         int64(appChannel.Quota),
				CreatedAt:     appChannel.CreatedAt.Unix(),
			},
		}, nil
	}
	return &appsvc.GetChannelByIdResp{}, nil
}
