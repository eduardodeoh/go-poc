package database

import (
	"testing"
)

func setupEnvDb(t testing.TB) {
	t.Setenv("DB_NAME", "test")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_USER", "test")
	t.Setenv("DB_PASSWORD", "test")
	t.Setenv("DB_SSLMODE", "disable")
	t.Setenv("DB_LOG_LEVEL", "debug")
}

func TestConfigDsn(t *testing.T) {
	setupEnvDb(t)

	type testCase struct {
		name        string
		envVar      string
		envVarValue string
		expected    string
	}

	dbTestCases := []testCase{
		{"should return Dsn value using DB_NAME envvar", "DB_NAME", "abc", "host=localhost port=5432 user=test password=test dbname=abc sslmode=disable"},
		{"should return Dsn value using DB_HOST envvar", "DB_HOST", "test", "host=test port=5432 user=test password=test dbname=test sslmode=disable"},
		{"should return Dsn value using DB_PORT envvar", "DB_PORT", "5433", "host=localhost port=5433 user=test password=test dbname=test sslmode=disable"},
		{"should return Dsn value using DB_USER envvar", "DB_USER", "jdoe", "host=localhost port=5432 user=jdoe password=test dbname=test sslmode=disable"},
		{"should return Dsn value using DB_PASSWORD envvar", "DB_PASSWORD", "pass", "host=localhost port=5432 user=test password=pass dbname=test sslmode=disable"},
		{"should return Dsn value using DB_SSLMODE envvar", "DB_SSLMODE", "enable", "host=localhost port=5432 user=test password=test dbname=test sslmode=enable"},
	}

	for _, tc := range dbTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envVar != "" {
				t.Setenv(tc.envVar, tc.envVarValue)
			}

			config, err := NewConfig()
			if err != nil {
				t.Fatalf("Error loading config: %v", err)
			}

			got := config.Dsn()

			if got != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, got)
			}
		})
	}

}
