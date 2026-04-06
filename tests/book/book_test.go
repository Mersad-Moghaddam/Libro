package book_test

import (
	"libro/apiSchema/bookSchema"
	"testing"
)

func TestBookmarkRequest(t *testing.T) {
	req := bookSchema.UpdateBookBookmarkRequest{CurrentPage: 10}
	if req.CurrentPage < 0 {
		t.Fatal("current page cannot be negative")
	}
}
