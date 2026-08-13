package main

import (
	"log/slog"
	"os"

	"github.com/sam-in07/E-Commerce_API_Go/internal/env"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING" ,  "host=localhost user=postgres password=admin dbname=ecom-api-go sslmode=disable"),
		},
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}

}
