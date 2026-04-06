package repositories

import (
	"gorm.io/gorm"
	"libro/models/purchaseLink"
)

type PurchaseLinkRepo struct{ db *gorm.DB }

func NewPurchaseLinkRepo(db *gorm.DB) *PurchaseLinkRepo               { return &PurchaseLinkRepo{db: db} }
func (r *PurchaseLinkRepo) Create(p *purchaseLink.PurchaseLink) error { return r.db.Create(p).Error }
func (r *PurchaseLinkRepo) FindByID(id uint) (*purchaseLink.PurchaseLink, error) {
	var p purchaseLink.PurchaseLink
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *PurchaseLinkRepo) Save(p *purchaseLink.PurchaseLink) error   { return r.db.Save(p).Error }
func (r *PurchaseLinkRepo) Delete(p *purchaseLink.PurchaseLink) error { return r.db.Delete(p).Error }
