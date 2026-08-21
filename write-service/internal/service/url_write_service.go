package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/deatil/go-encoding/encoding"

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
	salt := make([]byte, 8)
	rand.Read(salt) // Generate a random salt for uniqueness
	return fmt.Sprintf("localhost/%s", generateShortCode(originalURL, string(salt)))
}

func generateShortCode(originalUrl, salt string) string {
	ip := originalUrl + salt
	hash := sha256.Sum256([]byte(ip))
	encoded := encoding.FromBytes(hash[:]).Base62Encode().ToString()
	return encoded[:8]
}
