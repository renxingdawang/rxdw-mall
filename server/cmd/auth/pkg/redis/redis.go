package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type AuthRedisManager struct {
	client *redis.Client
}

func NewRedisManager(client *redis.Client) *AuthRedisManager {
	return &AuthRedisManager{client: client}
}

func (m *AuthRedisManager) StoreToken(ctx context.Context, UserId int32, token string, expiration time.Duration) error {
	key := fmt.Sprintf("token:%d", UserId)
	return m.client.Set(ctx, key, token, expiration).Err()
}
func (m *AuthRedisManager) GetToken(ctx context.Context, UserId int32) (string, error) {
	key := fmt.Sprintf("token:%d", UserId)
	token, err := m.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return " ", errors.New("token not found")
	} else if err != nil {
		return " ", err
	}
	return token, nil
}
