package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type InitialRepositories struct {
	UserRepo            *UserRepo
	BookRepo            *BookRepo
	WishlistRepo        *WishlistRepo
	PurchaseLinkRepo    *PurchaseLinkRepo
	ReadingProgressRepo *ReadingProgressRepo
	AuthRepo            *AuthRepo
}

func NewInitialRepositories(db *gorm.DB, rdb *redis.Client) *InitialRepositories {
	bookRepo := NewBookRepo(db)
	return &InitialRepositories{
		UserRepo:            NewUserRepo(db),
		BookRepo:            bookRepo,
		WishlistRepo:        NewWishlistRepo(db),
		PurchaseLinkRepo:    NewPurchaseLinkRepo(db),
		ReadingProgressRepo: NewReadingProgressRepo(bookRepo),
		AuthRepo:            NewAuthRepo(rdb),
	}
}
