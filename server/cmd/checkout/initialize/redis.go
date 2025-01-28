package initialize

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/renxingdawang/rxdw-mall/server/cmd/checkout/config"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GlobalServerConfig.RedisInfo.Host, config.GlobalServerConfig.RedisInfo.Port),
		Password: config.GlobalServerConfig.RedisInfo.Password,
		DB:       consts.RedisAuthClientDB,
	})
}
