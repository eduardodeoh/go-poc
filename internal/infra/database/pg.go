package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

// https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5#readme-adapters-for-3rd-party-loggers
type Logger struct {
	l *slog.Logger
}

func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {

	// Parsing database connection string
	config, err := pgxpool.ParseConfig(connString)

	if err != nil {
		return nil, fmt.Errorf("unable to parse pool config: %w", err)
	}

	// Database Connection Pool
	// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx#using-a-connection-pool
	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}

func NewPoolWithLogger(ctx context.Context, connString string, logger *slog.Logger, logLevel string) (*pgxpool.Pool, error) {
	// Parsing database connection string
	config, err := pgxpool.ParseConfig(connString)

	if err != nil {
		return nil, fmt.Errorf("unable to parse pool config: %w", err)
	}

	level, err := tracelog.LogLevelFromString(logLevel)
	if err != nil {
		return nil, fmt.Errorf("invalid database log level: %w", err)
	}

	// Update pool config
	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   &Logger{l: logger},
		LogLevel: level,
	}

	// Database Connection Pool
	// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx#using-a-connection-pool
	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}

func LogLevelFromString(level string) (tracelog.LogLevel, error) {
	l, err := tracelog.LogLevelFromString(level)
	if err != nil {
		return tracelog.LogLevelDebug, fmt.Errorf("database log level configuration: %w", err)
	}
	return l, nil
}

// https://github.com/mcosta74/pgx-slog/blob/main/adapter.go
func (logger *Logger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(data))
	for k, v := range data {
		attrs = append(attrs, slog.Any(k, v))
	}

	var lvl slog.Level
	switch level {
	case tracelog.LogLevelTrace:
		lvl = slog.LevelDebug - 1
		attrs = append(attrs, slog.Any("PGX_LOG_LEVEL", level))
	case tracelog.LogLevelDebug:
		lvl = slog.LevelDebug
	case tracelog.LogLevelInfo:
		lvl = slog.LevelInfo
	case tracelog.LogLevelWarn:
		lvl = slog.LevelWarn
	case tracelog.LogLevelError:
		lvl = slog.LevelError
	default:
		lvl = slog.LevelError
		attrs = append(attrs, slog.Any("INVALID_PGX_LOG_LEVEL", level))
	}

	logger.l.With("infra", "Postgresql").LogAttrs(context.Background(), lvl, msg, attrs...)
}
