package bookSchema

type CreateBookRequest struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	TotalPages int    `json:"totalPages"`
	Status     string `json:"status"`
}

type UpdateBookRequest struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	TotalPages int    `json:"totalPages"`
}

type UpdateBookStatusRequest struct {
	Status string `json:"status"`
}

type UpdateBookBookmarkRequest struct {
	CurrentPage int `json:"currentPage"`
}

type BookResponse struct {
	ID                 uint    `json:"id"`
	Title              string  `json:"title"`
	Author             string  `json:"author"`
	TotalPages         int     `json:"totalPages"`
	Status             string  `json:"status"`
	CurrentPage        int     `json:"currentPage"`
	RemainingPages     int     `json:"remainingPages"`
	ProgressPercentage float64 `json:"progressPercentage"`
	CompletedAt        *string `json:"completedAt"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
}

type BookListResponse struct {
	Items []BookResponse `json:"items"`
	Total int64          `json:"total"`
}

type BookDetailResponse struct {
	Item BookResponse `json:"item"`
}
