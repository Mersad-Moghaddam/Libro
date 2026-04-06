package core

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"libro/repositories"
	"libro/services/authService"
	"libro/services/bookService"
	"libro/services/readingService"
	"libro/services/wishlistService"
	"libro/statics/configs"
)

type Services struct {
	Auth     *authService.Service
	Book     *bookService.Service
	Reading  *readingService.Service
	Wishlist *wishlistService.Service
}

func InitServices(cfg *configs.Config, db *gorm.DB, rdb *redis.Client) (*Services, *repositories.InitialRepositories) {
	repos := repositories.NewInitialRepositories(db, rdb)
	return &Services{
		Auth:     authService.New(repos, cfg),
		Book:     bookService.New(repos),
		Reading:  readingService.New(repos),
		Wishlist: wishlistService.New(repos),
	}, repos
}
