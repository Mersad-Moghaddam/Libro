package reading_test

import (
	"libro/apiSchema/readingSchema"
	"testing"
)

func TestReadingProgressRequest(t *testing.T) {
	req := readingSchema.UpdateReadingProgressRequest{CurrentPage: 20}
	if req.CurrentPage < 0 {
		t.Fatal("invalid currentPage")
	}
}
