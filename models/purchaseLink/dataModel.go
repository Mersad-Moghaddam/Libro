package purchaseLink

import "time"

type PurchaseLink struct {
	ID         uint   `gorm:"primaryKey"`
	WishlistID uint   `gorm:"index;not null"`
	Label      string `gorm:"size:120;not null"`
	URL        string `gorm:"type:text;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (PurchaseLink) TableName() string { return "purchase_links" }
