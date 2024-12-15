package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/renxingdawang/rxdw-mall/server/cmd/auth/config"
	"github.com/renxingdawang/rxdw-mall/server/cmd/auth/initialize"
	mysql "github.com/renxingdawang/rxdw-mall/server/cmd/auth/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/cmd/auth/pkg/paseto"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	auth "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/auth/authservice"
	"net"
	"strconv"

	"log"
)

func main() {
	//IP,Port:=initialize.InitFlag()
	//r,info:=initialize.InitRegistry(Port)
	//initialization
	initialize.InitLogger()
	initialize.InitConfig()
	db := initialize.InitDB()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitRegistry(Port)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer func(p provider.OtelProvider, ctx context.Context) {
		err := p.Shutdown(ctx)
		if err != nil {

		}
	}(p, context.Background())

	tg, err := paseto.NewTokenGenerator(
		config.GlobalServerConfig.PasetoInfo.SecretKey,
		[]byte(config.GlobalServerConfig.PasetoInfo.Implicit))
	if err != nil {
		klog.Fatal(err)
	}

	svr := auth.NewServer(&AuthServiceImpl{
		AuthManger:     mysql.NewUserManager(db),
		TokenGenerator: tg,
	},
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}

}
