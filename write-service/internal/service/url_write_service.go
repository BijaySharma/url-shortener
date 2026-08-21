package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"

	"github.com/deatil/go-encoding/encoding"

	"github.com/BijaySharma/url-shortener/write-service/internal/repository"
)

type URLWriteService interface {
	// SaveURL saves the original URL and returns a unique short URL.
	SaveURL(ctx context.Context, originalURL string) (string, error)

	// GetOriginalURL retrieves the original URL for a given short URL.
	GetOriginalURL(ctx context.Context, shortCode string) (string, error)
}

type urlWriteService struct {
	url_repo repository.URLRepository
}

func NewURLWriteService(repo repository.URLRepository) URLWriteService {
	return &urlWriteService{url_repo: repo}
}

func (s *urlWriteService) SaveURL(ctx context.Context, originalURL string) (string, error) {
	shortCode := generateShortCode(originalURL)

	// Save the original URL and the generated short URL in the repository.
	return s.url_repo.SaveURL(ctx, originalURL, shortCode, nil)
}

func (s *urlWriteService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	return s.url_repo.GetOriginalURL(ctx, shortCode)
}

// v1: This function generates a unique short code for the given original URL using a hash function and random salt.
// v2: This uses redis global counter and sqids bijective function to generate short unique codes.

func generateShortCode(originalUrl string) string {
	// Approach 1: Use a hash function to generate a unique short code based on the original URL.
	salt := make([]byte, 8)
	rand.Read(salt) // Generate a random salt for uniqueness
	ip := originalUrl + string(salt)
	hash := sha256.Sum256([]byte(ip))
	encoded := encoding.FromBytes(hash[:]).Base62Encode().ToString()
	return encoded[:8]
}

func generateShortCode_v2() string {
	// TODO
	return ""
}
