package readingService

import (
	"time"

	"libro/apiSchema/readingSchema"
	"libro/repositories"
	"libro/statics/constants"
	"libro/statics/customErr"
)

type Service struct {
	repos *repositories.InitialRepositories
}

func New(repos *repositories.InitialRepositories) *Service { return &Service{repos: repos} }

func (s *Service) GetCurrentReadingBooks(userID uint) ([]readingSchema.ReadingProgressResponse, error) {
	items, err := s.repos.ReadingProgressRepo.ListCurrentReading(userID)
	if err != nil {
		return nil, err
	}
	result := make([]readingSchema.ReadingProgressResponse, 0, len(items))
	for _, b := range items {
		result = append(result, toResp(b.ID, b.CurrentPage, b.TotalPages, b.Status))
	}
	return result, nil
}

func (s *Service) UpdateReadingProgress(userID, bookID uint, currentPage int) (*readingSchema.ReadingProgressResponse, error) {
	b, err := s.repos.ReadingProgressRepo.FindBookByIDAndUser(bookID, userID)
	if err != nil {
		return nil, customErr.ErrNotFound
	}
	if currentPage < 0 || currentPage > b.TotalPages {
		return nil, customErr.ErrInvalidInput
	}
	b.CurrentPage = currentPage
	if currentPage == b.TotalPages {
		b.Status = constants.BookStatusFinished
		now := time.Now()
		b.CompletedAt = &now
	}
	if err := s.repos.ReadingProgressRepo.SaveBook(b); err != nil {
		return nil, err
	}
	resp := toResp(b.ID, b.CurrentPage, b.TotalPages, b.Status)
	return &resp, nil
}

func toResp(id uint, currentPage, totalPages int, status string) readingSchema.ReadingProgressResponse {
	remaining := totalPages - currentPage
	if remaining < 0 {
		remaining = 0
	}
	progress := 0.0
	if totalPages > 0 {
		progress = float64(currentPage) / float64(totalPages) * 100
	}
	return readingSchema.ReadingProgressResponse{BookID: id, CurrentPage: currentPage, TotalPages: totalPages, RemainingPages: remaining, ProgressPercentage: progress, Status: status}
}
