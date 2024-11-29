package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Name   string `mapstructure:"name" json:"name"`
	Host   string `mapstructure:"host" json:"host"`
	WsAddr string `mapstructure:"wsAddr" json:"wsAddr"`
}
