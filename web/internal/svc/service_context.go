package svc

import (
	"github.com/aheadIV/NightVoyager/accountsvc/client/accountlogsvc"
	"github.com/aheadIV/NightVoyager/accountsvc/client/accountsvc"
	"github.com/aheadIV/NightVoyager/web/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	AccountRPC    accountsvc.AccountSvc
	AccountLogRPC accountlogsvc.AccountLogSvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AccountRPC:    accountsvc.NewAccountSvc(zrpc.MustNewClient(c.AccountRPC)),
		AccountLogRPC: accountlogsvc.NewAccountLogSvc(zrpc.MustNewClient(c.AccountLogRPC)),
	}
}
