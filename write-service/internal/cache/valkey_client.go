package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
)

func NewClient(ctx context.Context, host string, port int, username, password *string) (*glide.Client, error) {
	cfg := config.NewClientConfiguration().
		WithAddress(&config.NodeAddress{Host: host, Port: port}).
		WithRequestTimeout(3 * time.Second)

	if username != nil && password != nil {
		cfg = cfg.WithCredentials(config.NewServerCredentials(*username, *password))

	}

	client, err := glide.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("create valkey client error: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if _, err := client.Ping(pingCtx); err != nil {
		client.Close()
		return nil, fmt.Errorf("ping valkey: %w", err)
	}

	log.Println("Valkey client is ready")
	return client, nil
}

func CloseClient(client *glide.Client) {
	log.Println("Closing valkey client")
	if client != nil {
		client.Close()
	}
}
