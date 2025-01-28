package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Salt     string `mapstructure:"salt" json:"salt"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}
type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	Prefix   string `mapstructure:"prefix" json:"prefix"`
}
type ServerConfig struct {
	Name           string           `mapstructure:"name" json:"name"`
	Host           string           `mapstructure:"host" json:"host"`
	MysqlInfo      MysqlConfig      `mapstructure:"mysql" json:"mysql"`
	RedisInfo      RedisConfig      `mapstructure:"redis" json:"redis"`
	OtelInfo       OtelConfig       `mapstructure:"otel" json:"otel"`
	CartSrvInfo    CartSrvConfig    `mapstructure:"cart_srv" json:"cart_srv"`
	OrderSrvInfo   OrderSrvConfig   `mapstructure:"order_srv" json:"order_srv"`
	PaymentSrvInfo PaymentSrvConfig `mapstructure:"payment_srv" json:"payment_srv"`
	ProductSrvInfo ProductSrvConfig `mapstructure:"Product_srv" json:"Product_srv"`
}
type CartSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
type PaymentSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
type ProductSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
