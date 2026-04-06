package initRepositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"libro/repositories"
)

func Init(db *gorm.DB, rdb *redis.Client) *repositories.InitialRepositories {
	return repositories.NewInitialRepositories(db, rdb)
}
