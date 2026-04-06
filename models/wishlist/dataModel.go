package wishlist

import (
	"libro/models/purchaseLink"
	"time"
)

type Wishlist struct {
	ID            uint                        `gorm:"primaryKey"`
	UserID        uint                        `gorm:"index;not null"`
	Title         string                      `gorm:"size:255;not null"`
	Author        string                      `gorm:"size:255;not null"`
	ExpectedPrice *float64                    `gorm:"type:decimal(10,2)"`
	Notes         *string                     `gorm:"type:text"`
	PurchaseLinks []purchaseLink.PurchaseLink `gorm:"foreignKey:WishlistID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Wishlist) TableName() string { return "wishlist" }
