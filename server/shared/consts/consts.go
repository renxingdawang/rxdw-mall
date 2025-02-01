package consts

import "time"

const (
	RxdwMall    = "rxdwMall"
	Issuer      = "FreeCar"
	Admin       = "Admin"
	User        = "User"
	ThirtyDays  = time.Hour * 24 * 30
	AccountID   = "accountID"
	ID          = "id"
	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName    = "port"
	PortFlagUsage   = "port"
	FreePortAddress = "localhost:0"
	CorsAddress     = "http://localhost:3000"
	TCP             = "tcp"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	HlogFilePath       = "./tmp/hlog/logs/"
	KlogFilePath       = "./tmp/klog/logs/"
	MySqlDSN           = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	RabbitMqURI        = "amqp://%s:%s@%s:%d/"
	MySQLImage         = "mysql:latest"
	MySQLContainerPort = "3306/tcp"
	MySQLContainerIP   = "121.40.228.214"
	MySQLPort          = "0"
	MySQLAdmin         = "root"
	DockerTestMySQLPwd = "123456"

	ApiConfigPath      = "./server/cmd/api/config.yaml"
	AuthConfigPath     = "./server/cmd/auth/config.yaml"
	CartConfigPath     = "./server/cmd/cart/config.yaml"
	CheckoutConfigPath = "./server/cmd/checkout/config.yaml"
	OrderConfigPath    = "./server/cmd/order/config.yaml"
	PaymentConfigPath  = "./server/cmd/payment/config.yaml"
	ProductConfigPath  = "./server/cmd/product/config.yaml"
	UserConfigPath     = "./server/cmd/user/config.yaml"

	RedisAuthClientDB = 0

	DelayQueue             = "payment_delay_queue"
	PaymentCancelledQueue  = "payment_cancelled_queue"
	OrderCancelFailedQueue = "order_cancel_failed_queue"
)

type OrderState string
type PayState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)
const (
	PayStateCreated  PayState = "created"
	PayStatePaid     PayState = "paid"
	PayStateCanceled PayState = "canceled"
)
