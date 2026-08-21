package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/BijaySharma/url-shortener/write-service/internal/db"
	"github.com/BijaySharma/url-shortener/write-service/internal/handler"
	"github.com/BijaySharma/url-shortener/write-service/internal/repository"
	"github.com/BijaySharma/url-shortener/write-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

var listen_port = ":8080"

func init() {
	if port := os.Getenv("PORT"); port != "" {
		listen_port = ":" + port
	}
}

func main() {
	db_pool := setupDatabase(context.Background())
	defer db.ClosePool(db_pool)

	mux := http.NewServeMux()
	registerHandlers(mux, db_pool)

	startServer(mux)
}

// dsnFromEnv builds postgres DSN from compose.yaml env vars.
func dsnFromEnv() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return db.GetDsn(host, port, user, password, dbname)
}

// setupDatabase runs migrations then opens the app's connection pool.
func setupDatabase(ctx context.Context) *pgxpool.Pool {
	dsn := dsnFromEnv()

	if err := db.RunMigrations(dsn, "file://migrations"); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	pool, err := db.NewPool(ctx, dsn)
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	return pool
}

func registerHandlers(mux *http.ServeMux, dbPool *pgxpool.Pool) {
	mux.HandleFunc("/health", handler.HealthCheckHandler)

	// Register the URL write handler
	url_repo := repository.NewURLRepository(dbPool)
	url_svc := service.NewURLWriteService(url_repo)
	urlServiceHandlers := handler.NewURLServiceHandlers(url_svc)
	mux.HandleFunc("POST /urls", urlServiceHandlers.SaveURLHandler)
	mux.HandleFunc("GET /urls", urlServiceHandlers.GetOriginalURLHandler)
}

func startServer(mux *http.ServeMux) {
	log.Printf("Starting server on port %s", listen_port)
	log.Fatal(http.ListenAndServe(listen_port, mux))
}
