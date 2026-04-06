package wishlist_test

import (
	"libro/apiSchema/purchaseLinkSchema"
	"testing"
)

func TestPurchaseLinkRequest(t *testing.T) {
	req := purchaseLinkSchema.CreatePurchaseLinkRequest{Label: "Store", URL: "https://example.com"}
	if req.URL == "" || req.Label == "" {
		t.Fatal("invalid purchase link")
	}
}
