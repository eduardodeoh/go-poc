package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/eduardodeoh/go-poc/internal/core/config"
	"github.com/eduardodeoh/go-poc/internal/infra/database"
)

func main() {
	var appEnv string = os.Getenv("APP_ENV")

	// Initialize Logger
	logger := initializeLogger(appEnv)

	// Initialize Config
	appConfig, err := config.NewConfig()

	if err != nil {
		logger.Error("error loading config", "details", err)
		os.Exit(1)
	}
	logger.Info("App Config loaded!")

	// Initialize Database Pool
	databaseLogger := database.NewLogger(logger)
	databaseLogLevel, err := database.LogLevelFromString(appConfig.Db.LogLevel)
	if err != nil {
		logger.Error("error parsing database log level", "details", err)
		os.Exit(1)
	}

	logger.Info("Database log level", "value", appConfig.Db.LogLevel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.NewPool(ctx, appConfig.DbDsn(), databaseLogger, databaseLogLevel)

	if err != nil {
		logger.Error("Database failed", "details", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("Database loaded successfully")

	logger.Info(
		"App loaded successfully",
		slog.String("appEnv", appEnv),
	)
}

func initializeLogger(appEnv string) *slog.Logger {
	var logLevel = new(slog.LevelVar) // Info by default

	if appEnv == "development" {
		logLevel.Set(slog.LevelDebug)
	}

	handlerOpts := &slog.HandlerOptions{Level: logLevel}
	var loggerHandler slog.Handler = slog.NewTextHandler(os.Stdout, handlerOpts)
	if appEnv == "production" {
		loggerHandler = slog.NewJSONHandler(os.Stdout, handlerOpts)
	}
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)

	logger.Info(
		"Logger loaded!",
		slog.String("defaultLogLevel", logLevel.String()),
	)

	return logger
}
