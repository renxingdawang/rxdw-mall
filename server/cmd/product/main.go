package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/renxingdawang/rxdw-mall/server/cmd/product/config"
	"github.com/renxingdawang/rxdw-mall/server/cmd/product/initialize"
	"github.com/renxingdawang/rxdw-mall/server/cmd/product/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/cmd/product/pkg/redis"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	product "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/product/productcatalogservice"
	"net"
	"strconv"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	fmt.Println("success")
	cache := initialize.InitRedis()
	db := initialize.InitDB()
	IP, Port := initialize.InitFlag()
	fmt.Println("flag ok")
	r, info := initialize.InitRegistry(Port)
	fmt.Println("register ok")
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
	fmt.Println("ok tg")
	productMysqlManager := mysql.NewProductMysqlManager(db)
	categoryMysqlManager := mysql.NewCategoryMysqlManager(db)
	productRedisManager := redis.NewRedisManager(productMysqlManager, cache)
	srv := product.NewServer(&ProductCatalogServiceImpl{
		ProductMysqlManager:  productMysqlManager,
		CategoryMysqlManager: categoryMysqlManager,
		ProductRedisManager:  productRedisManager,
	},
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)
	fmt.Println("success all")
	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
