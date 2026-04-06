package readingSchema

type UpdateReadingProgressRequest struct {
	CurrentPage int `json:"currentPage"`
}

type ReadingProgressResponse struct {
	BookID             uint    `json:"bookId"`
	CurrentPage        int     `json:"currentPage"`
	TotalPages         int     `json:"totalPages"`
	RemainingPages     int     `json:"remainingPages"`
	ProgressPercentage float64 `json:"progressPercentage"`
	Status             string  `json:"status"`
}
