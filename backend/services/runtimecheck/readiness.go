package runtimecheck

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service struct {
	db    *gorm.DB
	redis *redis.Client
}

func New(db *gorm.DB, redis *redis.Client) *Service {
	return &Service{db: db, redis: redis}
}

func (s *Service) Check(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("mysql handle: %w", err)
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("mysql ping: %w", err)
	}
	if err := s.redis.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis ping: %w", err)
	}
	return nil
}
