package repositories

import (
	"gorm.io/gorm"
	"libro/models/wishlist"
)

type WishlistRepo struct{ db *gorm.DB }

func NewWishlistRepo(db *gorm.DB) *WishlistRepo           { return &WishlistRepo{db: db} }
func (r *WishlistRepo) Create(w *wishlist.Wishlist) error { return r.db.Create(w).Error }
func (r *WishlistRepo) CountByUser(userID uint) (int64, error) {
	var t int64
	return t, r.db.Model(&wishlist.Wishlist{}).Where("user_id = ?", userID).Count(&t).Error
}
func (r *WishlistRepo) ListByUser(userID uint, limit, offset int) ([]wishlist.Wishlist, error) {
	var l []wishlist.Wishlist
	err := r.db.Preload("PurchaseLinks").Where("user_id = ?", userID).Limit(limit).Offset(offset).Order("id desc").Find(&l).Error
	return l, err
}
func (r *WishlistRepo) FindByIDAndUser(id, userID uint) (*wishlist.Wishlist, error) {
	var w wishlist.Wishlist
	err := r.db.Preload("PurchaseLinks").Where("id = ? and user_id = ?", id, userID).First(&w).Error
	if err != nil {
		return nil, err
	}
	return &w, nil
}
func (r *WishlistRepo) Save(w *wishlist.Wishlist) error   { return r.db.Save(w).Error }
func (r *WishlistRepo) Delete(w *wishlist.Wishlist) error { return r.db.Delete(w).Error }
