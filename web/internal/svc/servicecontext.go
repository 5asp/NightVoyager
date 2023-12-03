package svc

import (
	"github.com/aheadIV/NightVoyager/appsvc/appsvcclient"
	"github.com/aheadIV/NightVoyager/web/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	App    appsvcclient.Appsvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		App:    appsvcclient.NewAppsvc(zrpc.MustNewClient(c.AppRpc)),
	}
}
