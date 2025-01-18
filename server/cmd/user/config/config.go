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
type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Host      string      `mapstructure:"host" json:"host"`
	WsAddr    string      `mapstructure:"wsAddr" json:"wsAddr"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	OtelInfo  OtelConfig  `mapstructure:"otel" json:"otel"`
}
type AuthSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
