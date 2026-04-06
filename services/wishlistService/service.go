package wishlistService

import (
	"net/url"
	"time"

	"libro/apiSchema/purchaseLinkSchema"
	"libro/apiSchema/wishlistSchema"
	"libro/models/commonPagination"
	"libro/models/purchaseLink"
	"libro/models/wishlist"
	"libro/repositories"
	"libro/statics/customErr"
)

type Service struct {
	repos *repositories.InitialRepositories
}

func New(repos *repositories.InitialRepositories) *Service { return &Service{repos: repos} }

func (s *Service) CreateWishlistItem(userID uint, req wishlistSchema.CreateWishlistRequest) (*wishlistSchema.WishlistResponse, error) {
	w := &wishlist.Wishlist{UserID: userID, Title: req.Title, Author: req.Author, ExpectedPrice: req.ExpectedPrice, Notes: req.Notes}
	if err := s.repos.WishlistRepo.Create(w); err != nil {
		return nil, err
	}
	resp := toWishlistResponse(*w)
	return &resp, nil
}
func (s *Service) GetWishlist(userID uint, req commonPagination.PageRequest) (*wishlistSchema.WishlistListResponse, error) {
	total, err := s.repos.WishlistRepo.CountByUser(userID)
	if err != nil {
		return nil, err
	}
	items, err := s.repos.WishlistRepo.ListByUser(userID, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, err
	}
	resp := make([]wishlistSchema.WishlistResponse, 0, len(items))
	for _, w := range items {
		resp = append(resp, toWishlistResponse(w))
	}
	return &wishlistSchema.WishlistListResponse{Items: resp, Total: total}, nil
}
func (s *Service) GetWishlistItemByID(userID, id uint) (*wishlistSchema.WishlistResponse, error) {
	w, err := s.repos.WishlistRepo.FindByIDAndUser(id, userID)
	if err != nil {
		return nil, customErr.ErrNotFound
	}
	resp := toWishlistResponse(*w)
	return &resp, nil
}
func (s *Service) UpdateWishlistItem(userID, id uint, req wishlistSchema.UpdateWishlistRequest) (*wishlistSchema.WishlistResponse, error) {
	w, err := s.repos.WishlistRepo.FindByIDAndUser(id, userID)
	if err != nil {
		return nil, customErr.ErrNotFound
	}
	w.Title = req.Title
	w.Author = req.Author
	w.ExpectedPrice = req.ExpectedPrice
	w.Notes = req.Notes
	if err := s.repos.WishlistRepo.Save(w); err != nil {
		return nil, err
	}
	resp := toWishlistResponse(*w)
	return &resp, nil
}
func (s *Service) DeleteWishlistItem(userID, id uint) error {
	w, err := s.repos.WishlistRepo.FindByIDAndUser(id, userID)
	if err != nil {
		return customErr.ErrNotFound
	}
	return s.repos.WishlistRepo.Delete(w)
}

func (s *Service) AddPurchaseLink(userID, wishlistID uint, req purchaseLinkSchema.CreatePurchaseLinkRequest) (*purchaseLinkSchema.PurchaseLinkResponse, error) {
	if !validURL(req.URL) {
		return nil, customErr.ErrInvalidInput
	}
	w, err := s.repos.WishlistRepo.FindByIDAndUser(wishlistID, userID)
	if err != nil {
		return nil, customErr.ErrNotFound
	}
	p := &purchaseLink.PurchaseLink{WishlistID: w.ID, Label: req.Label, URL: req.URL}
	if err := s.repos.PurchaseLinkRepo.Create(p); err != nil {
		return nil, err
	}
	resp := toPurchaseLinkResponse(*p)
	return &resp, nil
}
func (s *Service) UpdatePurchaseLink(userID, wishlistID, linkID uint, req purchaseLinkSchema.UpdatePurchaseLinkRequest) (*purchaseLinkSchema.PurchaseLinkResponse, error) {
	if !validURL(req.URL) {
		return nil, customErr.ErrInvalidInput
	}
	if _, err := s.repos.WishlistRepo.FindByIDAndUser(wishlistID, userID); err != nil {
		return nil, customErr.ErrNotFound
	}
	p, err := s.repos.PurchaseLinkRepo.FindByID(linkID)
	if err != nil || p.WishlistID != wishlistID {
		return nil, customErr.ErrNotFound
	}
	p.Label = req.Label
	p.URL = req.URL
	if err := s.repos.PurchaseLinkRepo.Save(p); err != nil {
		return nil, err
	}
	resp := toPurchaseLinkResponse(*p)
	return &resp, nil
}
func (s *Service) DeletePurchaseLink(userID, wishlistID, linkID uint) error {
	if _, err := s.repos.WishlistRepo.FindByIDAndUser(wishlistID, userID); err != nil {
		return customErr.ErrNotFound
	}
	p, err := s.repos.PurchaseLinkRepo.FindByID(linkID)
	if err != nil || p.WishlistID != wishlistID {
		return customErr.ErrNotFound
	}
	return s.repos.PurchaseLinkRepo.Delete(p)
}

func validURL(v string) bool { _, err := url.ParseRequestURI(v); return err == nil }
func toPurchaseLinkResponse(p purchaseLink.PurchaseLink) purchaseLinkSchema.PurchaseLinkResponse {
	return purchaseLinkSchema.PurchaseLinkResponse{ID: p.ID, Label: p.Label, URL: p.URL, CreatedAt: p.CreatedAt.Format(time.RFC3339), UpdatedAt: p.UpdatedAt.Format(time.RFC3339)}
}
func toWishlistResponse(w wishlist.Wishlist) wishlistSchema.WishlistResponse {
	links := make([]purchaseLinkSchema.PurchaseLinkResponse, 0, len(w.PurchaseLinks))
	for _, l := range w.PurchaseLinks {
		links = append(links, toPurchaseLinkResponse(l))
	}
	return wishlistSchema.WishlistResponse{ID: w.ID, Title: w.Title, Author: w.Author, ExpectedPrice: w.ExpectedPrice, Notes: w.Notes, PurchaseLinks: links, CreatedAt: w.CreatedAt.Format(time.RFC3339), UpdatedAt: w.UpdatedAt.Format(time.RFC3339)}
}
