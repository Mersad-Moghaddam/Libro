package wishlistSchema

import "libro/apiSchema/purchaseLinkSchema"

type CreateWishlistRequest struct {
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	ExpectedPrice *float64 `json:"expectedPrice"`
	Notes         *string  `json:"notes"`
}

type UpdateWishlistRequest = CreateWishlistRequest

type WishlistResponse struct {
	ID            uint                                      `json:"id"`
	Title         string                                    `json:"title"`
	Author        string                                    `json:"author"`
	ExpectedPrice *float64                                  `json:"expectedPrice"`
	Notes         *string                                   `json:"notes"`
	PurchaseLinks []purchaseLinkSchema.PurchaseLinkResponse `json:"purchaseLinks"`
	CreatedAt     string                                    `json:"createdAt"`
	UpdatedAt     string                                    `json:"updatedAt"`
}

type WishlistListResponse struct {
	Items []WishlistResponse `json:"items"`
	Total int64              `json:"total"`
}
