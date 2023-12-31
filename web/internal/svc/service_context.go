package svc

import (
	"github.com/aheadIV/NightVoyager/accountsvc/client/accountlogsvc"
	"github.com/aheadIV/NightVoyager/accountsvc/client/accountsvc"
	"github.com/aheadIV/NightVoyager/appsvc/client/appsvc"
	"github.com/aheadIV/NightVoyager/sendsvc/client/tplsendsvc"
	"github.com/aheadIV/NightVoyager/sendsvc/client/tplsignsvc"
	"github.com/aheadIV/NightVoyager/sendsvc/client/tplsvc"
	"github.com/aheadIV/NightVoyager/web/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	AccountRPC    accountsvc.AccountSvc
	AccountLogRPC accountlogsvc.AccountLogSvc
	AppRPC        appsvc.Appsvc
	TplSendRPC    tplsendsvc.TplSendSvc
	TplSignRPC    tplsignsvc.TplSignSvc
	TplRPC        tplsvc.TplSvc
	ClientIP      string
	Device        string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AccountRPC:    accountsvc.NewAccountSvc(zrpc.MustNewClient(c.AccountRPC)),
		AccountLogRPC: accountlogsvc.NewAccountLogSvc(zrpc.MustNewClient(c.AccountLogRPC)),
		AppRPC:        appsvc.NewAppsvc(zrpc.MustNewClient(c.AppRPC)),
		TplSendRPC:    tplsendsvc.NewTplSendSvc(zrpc.MustNewClient(c.TplSendRPC)),
		TplSignRPC:    tplsignsvc.NewTplSignSvc(zrpc.MustNewClient(c.TplSignRPC)),
		TplRPC:        tplsvc.NewTplSvc(zrpc.MustNewClient(c.TplRPC)),
	}
}
