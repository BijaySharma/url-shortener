package repository

import "github.com/jackc/pgx/v5/pgxpool"

type URLRepository interface {
	// SaveURL saves the original URL and returns a unique short URL.
	SaveURL(originalURL string) (string, error)

	// GetOriginalURL retrieves the original URL for a given short URL.
	GetOriginalURL(shortURL string) (string, error)
}

type pgURLRepository struct {
	db_pool *pgxpool.Pool
}

func (r *pgURLRepository) SaveURL(originalURL string) (string, error) {
	// Implementation for saving the original URL and returning a unique short URL.
	return "", nil
}

func (r *pgURLRepository) GetOriginalURL(shortURL string) (string, error) {
	// Implementation for retrieving the original URL for a given short URL.
	return "", nil
}
