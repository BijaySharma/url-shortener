package service

import (
	"context"

	"github.com/BijaySharma/url-shortener/write-service/internal/repository"
)

type URLWriteService interface {
	// SaveURL saves the original URL and returns a unique short URL.
	SaveURL(ctx context.Context, originalURL string) (string, error)

	// GetOriginalURL retrieves the original URL for a given short URL.
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}

type urlWriteService struct {
	url_repo repository.URLRepository
}

func NewURLWriteService(repo repository.URLRepository) URLWriteService {
	return &urlWriteService{url_repo: repo}
}

func (s *urlWriteService) SaveURL(ctx context.Context, originalURL string) (string, error) {
	shortURL := generateShortURL(originalURL)

	// Save the original URL and the generated short URL in the repository.
	return s.url_repo.SaveURL(ctx, originalURL, shortURL, nil)
}

func (s *urlWriteService) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	return s.url_repo.GetOriginalURL(ctx, shortURL)
}

func generateShortURL(originalURL string) string {

	return "short_" + originalURL // TODO: Replace with actual short URL generation logic.
}
