package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository interface {
	// SaveURL saves the original URL and returns a unique short URL.
	SaveURL(ctx context.Context, originalURL, shortURL string, expirationTime *time.Time) (string, error)

	// GetOriginalURL retrieves the original URL for a given short URL.
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}

type pgURLRepository struct {
	db_pool *pgxpool.Pool
}

func NewURLRepository(db_pool *pgxpool.Pool) URLRepository {
	return &pgURLRepository{db_pool: db_pool}
}

var ErrURLNotFound = errors.New("url not found")

func (r *pgURLRepository) SaveURL(ctx context.Context, url, shortURL string, expirationTime *time.Time) (string, error) {
	// Expiration time is optional; if nil, the URL will not expire.
	query := `INSERT INTO urls (original_url, short_url, expiration_time) VALUES ($1, $2, $3) RETURNING short_url`
	var returnedShortURL string
	err := r.db_pool.QueryRow(ctx, query, url, shortURL, expirationTime).Scan(&returnedShortURL)
	if err != nil {
		return "", err
	}
	return returnedShortURL, nil
}

func (r *pgURLRepository) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	query := `SELECT original_url FROM urls WHERE short_url = $1`
	var originalURL string
	err := r.db_pool.QueryRow(ctx, query, shortURL).Scan(&originalURL)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", ErrURLNotFound
		}
		return "", err
	}
	return originalURL, nil
}
