package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, dsn_or_url string) (*pgxpool.Pool, error) {
	// Parsing database connection string
	config, err := pgxpool.ParseConfig(dsn_or_url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse pool config: %v\n", err)
		os.Exit(1)
	}

	// Database Connection Pool
	// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx#using-a-connection-pool
	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	if err := pool.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v, connection string: %s\n", err, dsn_or_url)
		os.Exit(1)
	}

	return pool, nil
}
