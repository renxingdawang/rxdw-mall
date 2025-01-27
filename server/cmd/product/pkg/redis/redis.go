package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/renxingdawang/rxdw-mall/server/cmd/product/config"

	"github.com/renxingdawang/rxdw-mall/server/cmd/product/pkg/mysql"
	"time"
)

type ProductRedisManager struct {
	client              *redis.Client
	prefix              string
	productMysqlManager mysql.ProductMysqlManager
}

func NewRedisManager(p mysql.ProductMysqlManager, client *redis.Client) *ProductRedisManager {
	return &ProductRedisManager{productMysqlManager: p, client: client, prefix: config.GlobalServerConfig.RedisInfo.Prefix}
}

func (p *ProductRedisManager) GetByID(ctx context.Context, productId int) (product mysql.Product, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d", p.prefix, "product_by_id", productId)
	cacheResult := p.client.Get(ctx, cacheKey)
	err = func() error {
		err1 := cacheResult.Err()
		if err1 != nil {
			return err1
		}
		cacheResultByte, err2 := cacheResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cacheResultByte, &product)
		if err3 != nil {
			return err3
		}
		return nil
	}()
	if err != nil {
		product, err = p.productMysqlManager.GetById(ctx, productId)
		if err != nil {
			return mysql.Product{}, err
		}
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = p.client.Set(ctx, cacheKey, encoded, time.Hour)
	}
	return
}
