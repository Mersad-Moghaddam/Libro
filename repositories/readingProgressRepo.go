package repositories

import "libro/models/book"

type ReadingProgressRepo struct{ bookRepo *BookRepo }

func NewReadingProgressRepo(bookRepo *BookRepo) *ReadingProgressRepo {
	return &ReadingProgressRepo{bookRepo: bookRepo}
}
func (r *ReadingProgressRepo) ListCurrentReading(userID uint) ([]book.Book, error) {
	return r.bookRepo.ListCurrentReading(userID)
}
func (r *ReadingProgressRepo) FindBookByIDAndUser(id, userID uint) (*book.Book, error) {
	return r.bookRepo.FindByIDAndUser(id, userID)
}
func (r *ReadingProgressRepo) SaveBook(b *book.Book) error { return r.bookRepo.Save(b) }
