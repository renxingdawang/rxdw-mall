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
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/config"
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/initialize"
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/pkg/mq"
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	payment "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/payment/paymentservice"
	"net"
	"strconv"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	db := initialize.InitDB()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitRegistry(Port)
	orderClient := initialize.InitOrder()
	mqConn := initialize.InitMQ()
	rabbitMQ := mq.NewRabbitMQ(mqConn)
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
	paymentService := &PaymentServiceImpl{
		PaymentLogMysqlManager: mysql.NewPaymentLogMysqlManager(db),
		OrderManager:           orderClient,
		RabbitMQ:               rabbitMQ,
	}
	srv := payment.NewServer(
		paymentService,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)
	go StartPaymentCancelConsumer(rabbitMQ, paymentService)
	go StartPaymentCompensationConsumer(rabbitMQ, paymentService)
	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
