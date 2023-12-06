package main

import (
	"flag"
	"fmt"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/config"
	accountlogsvcServer "github.com/aheadIV/NightVoyager/accountsvc/internal/server/accountlogsvc"
	accountsvcServer "github.com/aheadIV/NightVoyager/accountsvc/internal/server/accountsvc"
	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/accountsvc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		accountsvc.RegisterAccountSvcServer(grpcServer, accountsvcServer.NewAccountSvcServer(ctx))
		accountsvc.RegisterAccountLogSvcServer(grpcServer, accountlogsvcServer.NewAccountLogSvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
