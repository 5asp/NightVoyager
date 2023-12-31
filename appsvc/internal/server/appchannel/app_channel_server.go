// Code generated by goctl. DO NOT EDIT.
// Source: appsvc.proto

package server

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/logic/appchannel"
	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
)

type AppChannelServer struct {
	svcCtx *svc.ServiceContext
	appsvc.UnimplementedAppChannelServer
}

func NewAppChannelServer(svcCtx *svc.ServiceContext) *AppChannelServer {
	return &AppChannelServer{
		svcCtx: svcCtx,
	}
}

func (s *AppChannelServer) CreateChannel(ctx context.Context, in *appsvc.CreateChannelReq) (*appsvc.CreateChannelResp, error) {
	l := appchannellogic.NewCreateChannelLogic(ctx, s.svcCtx)
	return l.CreateChannel(in)
}

func (s *AppChannelServer) GetChannelById(ctx context.Context, in *appsvc.GetChannelByIdReq) (*appsvc.GetChannelByIdResp, error) {
	l := appchannellogic.NewGetChannelByIdLogic(ctx, s.svcCtx)
	return l.GetChannelById(in)
}

func (s *AppChannelServer) ChannelList(ctx context.Context, in *appsvc.ChannelListReq) (*appsvc.ChannelListResp, error) {
	l := appchannellogic.NewChannelListLogic(ctx, s.svcCtx)
	return l.ChannelList(in)
}
