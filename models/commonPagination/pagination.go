package commonPagination

type PageRequest struct {
	Page  int
	Limit int
}

type PageResult struct {
	Page   int   `json:"page"`
	Limit  int   `json:"limit"`
	Total  int64 `json:"total"`
	Offset int   `json:"offset"`
}
