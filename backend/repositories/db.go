package repositories

import (
	"fmt"

	"gorm.io/gorm"
	"libro-backend/models/book"
	"libro-backend/models/purchaseLink"
	"libro-backend/models/user"
	"libro-backend/models/wishlist"
)

func AssertSchema(db *gorm.DB) error {
	checks := []struct {
		model   any
		columns []string
	}{
		{&user.User{}, []string{"id", "name", "email", "password_hash", "reminder_enabled", "reminder_time", "reminder_frequency", "created_at", "updated_at"}},
		{&book.Book{}, []string{"id", "user_id", "title", "author", "total_pages", "status", "current_page", "completed_at", "created_at", "updated_at"}},
		{&wishlist.Wishlist{}, []string{"id", "user_id", "title", "author", "expected_price", "notes", "created_at", "updated_at"}},
		{&purchaseLink.PurchaseLink{}, []string{"id", "wishlist_id", "label", "alias", "url", "created_at", "updated_at"}},
	}

	for _, check := range checks {
		if !db.Migrator().HasTable(check.model) {
			return fmt.Errorf("missing table for model %T: run SQL migrations", check.model)
		}
		for _, column := range check.columns {
			if !db.Migrator().HasColumn(check.model, column) {
				return fmt.Errorf("missing column %q for model %T: run SQL migrations", column, check.model)
			}
		}
	}
	return nil
}
