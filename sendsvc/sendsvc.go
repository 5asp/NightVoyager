package main

import (
	"flag"
	"fmt"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/config"
	tplsendsvcServer "github.com/aheadIV/NightVoyager/sendsvc/internal/server/tplsendsvc"
	tplsignsvcServer "github.com/aheadIV/NightVoyager/sendsvc/internal/server/tplsignsvc"
	tplsvcServer "github.com/aheadIV/NightVoyager/sendsvc/internal/server/tplsvc"
	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sendsvc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sendsvc.RegisterTplSvcServer(grpcServer, tplsvcServer.NewTplSvcServer(ctx))
		sendsvc.RegisterTplSignSvcServer(grpcServer, tplsignsvcServer.NewTplSignSvcServer(ctx))
		sendsvc.RegisterTplSendSvcServer(grpcServer, tplsendsvcServer.NewTplSendSvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
