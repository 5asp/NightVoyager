package main

import (
	"flag"
	"fmt"

	"github.com/aheadIV/NightVoyager/appsvc/internal/config"
	appchannelServer "github.com/aheadIV/NightVoyager/appsvc/internal/server/appchannel"
	appsvcServer "github.com/aheadIV/NightVoyager/appsvc/internal/server/appsvc"
	sendlogsvcServer "github.com/aheadIV/NightVoyager/appsvc/internal/server/sendlogsvc"
	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/appsvc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		appsvc.RegisterAppsvcServer(grpcServer, appsvcServer.NewAppsvcServer(ctx))
		appsvc.RegisterSendLogsvcServer(grpcServer, sendlogsvcServer.NewSendLogsvcServer(ctx))
		appsvc.RegisterAppChannelServer(grpcServer, appchannelServer.NewAppChannelServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
