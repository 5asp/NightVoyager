package svc

import (
	"github.com/aheadIV/NightVoyager/appsvc/appsvcclient"
	"github.com/aheadIV/NightVoyager/send/internal/config"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AppRPC   appsvcclient.Appsvc
	JsStream jetstream.JetStream
}

func NewServiceContext(c config.Config) *ServiceContext {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := jetstream.New(nc)

	return &ServiceContext{
		Config:   c,
		AppRPC:   appsvcclient.NewAppsvc(zrpc.MustNewClient(c.AppRPC)),
		JsStream: js,
	}
}
