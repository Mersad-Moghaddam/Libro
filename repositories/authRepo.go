package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"libro/statics/constants"
)

type AuthRepo struct{ redis *redis.Client }

func NewAuthRepo(r *redis.Client) *AuthRepo { return &AuthRepo{redis: r} }
func (r *AuthRepo) StoreRefreshToken(userID uint, token string, ttl time.Duration) error {
	return r.redis.Set(context.Background(), fmt.Sprintf("%s%d", constants.RedisRefreshPrefix, userID), token, ttl).Err()
}
func (r *AuthRepo) GetRefreshToken(userID uint) (string, error) {
	return r.redis.Get(context.Background(), fmt.Sprintf("%s%d", constants.RedisRefreshPrefix, userID)).Result()
}
func (r *AuthRepo) DeleteRefreshToken(userID uint) error {
	return r.redis.Del(context.Background(), fmt.Sprintf("%s%d", constants.RedisRefreshPrefix, userID)).Err()
}
