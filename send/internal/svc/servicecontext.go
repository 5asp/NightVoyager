package svc

import (
	"github.com/aheadIV/NightVoyager/appsvc/appsvcclient"
	"github.com/aheadIV/NightVoyager/send/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	AppRPC appsvcclient.Appsvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AppRPC: appsvcclient.NewAppsvc(zrpc.MustNewClient(c.AppRPC)),
	}
}
