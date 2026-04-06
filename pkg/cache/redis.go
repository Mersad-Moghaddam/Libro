package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"libro/statics/configs"
)

func InitRedis(cfg *configs.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr(), Password: cfg.RedisPassword})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
