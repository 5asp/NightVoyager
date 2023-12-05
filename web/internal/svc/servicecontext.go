package svc

import (
	"github.com/aheadIV/NightVoyager/appsvc/appsvcclient"
	"github.com/aheadIV/NightVoyager/web/internal/config"
	"github.com/aheadIV/NightVoyager/web/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	App               appsvcclient.Appsvc
	WebAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		App:               appsvcclient.NewAppsvc(zrpc.MustNewClient(c.AppRpc)),
		WebAuthMiddleware: middleware.NewWebAuthMiddleware().Handle,
	}
}
