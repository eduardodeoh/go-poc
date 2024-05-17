package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/eduardodeoh/go-poc/internal/core/config"
	"github.com/eduardodeoh/go-poc/internal/infra/postgresql"
)

func main() {
	// Initialize Config
	appConfig, err := config.New()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize App Config: %v\n", err)
		os.Exit(1)
	}

	// Initialize Database Pool
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := postgresql.NewPool(ctx, appConfig.DbDsn())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize database pool: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	fmt.Println("Database connected successfully!")
}
