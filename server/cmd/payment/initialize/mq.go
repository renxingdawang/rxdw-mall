package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/config"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"github.com/streadway/amqp"
)

func InitMQ() *amqp.Connection {
	c := config.GlobalServerConfig.RabbitMqInfo
	amqpConn, err := amqp.Dial(fmt.Sprintf(consts.RabbitMqURI, c.User, c.Password, c.Host, c.Port))
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	return amqpConn
}
