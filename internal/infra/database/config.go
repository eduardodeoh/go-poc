package database

import (
	"fmt"

	"github.com/ardanlabs/conf/v3"
)

type Db struct {
	Url      string `conf:"env:DB_URL"`
	Name     string `conf:"env:DB_NAME,required"`
	Host     string `conf:"env:DB_HOST,default:localhost"`
	Port     int    `conf:"env:DB_PORT,default:5432"`
	User     string `conf:"env:DB_USER,default:postgres"`
	Password string `conf:"env:DB_PASSWORD,default:postgres"`
	SSLMode  string `conf:"env:DB_SSLMODE,default:disable"`
	// Valid log levels: https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5/tracelog#LogLevelFromString
	LogLevel string `conf:"env:DB_LOG_LEVEL,default:info"`
}

type Pool struct {
	// https://pkg.go.dev/github.com/jackc/pgx#ConnConfig
	//https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5/pgxpool#Config
	MaxConn               int    `conf:"env:DB_POOL_MAX_CONN,default:5"`
	MinConn               int    `conf:"env:DB_POOL_MIN_CONN,default:0"`
	MaxConnLifetime       string `conf:"env:DB_POOL_MAX_CONN_LIFETIME,default:5m"`
	MaxConnIdleTime       string `conf:"env:DB_POOL_MAX_CONN_IDLE_TIME,default:5m"`
	HealthCheckPeriod     string `conf:"env:DB_POOL_HEALTH_CHECK_PERIOD,default:5m"`
	MaxConnLifetimeJitter string `conf:"env:DB_POOL_MAX_CONN_LIFETIME_JITTER,default:30s"`
}

type Config struct {
	Db
	Pool
}

func NewConfig() (*Config, error) {
	c := Config{}
	if _, err := conf.Parse("", &c); err != nil {
		return &Config{}, fmt.Errorf("fail to parse config: %w", err)
	}

	return &c, nil
}

func (c *Config) Dsn() string {
	// https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
	base := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	return fmt.Sprintf(base, c.Db.Host, c.Db.Port, c.Db.User, c.Db.Password, c.Db.Name, c.Db.SSLMode)
}

func (c *Config) DsnWithPoolOptions() string {
	// https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5/pgxpool#ParseConfig
	base_dsn := c.Dsn()
	pool := "pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s pool_health_check_period=%s pool_max_conn_lifetime_jitter=%s"
	pool_dsn := fmt.Sprintf(pool, c.Pool.MaxConn, c.Pool.MinConn, c.Pool.MaxConnLifetime, c.Pool.MaxConnIdleTime, c.Pool.HealthCheckPeriod, c.Pool.MaxConnLifetimeJitter)

	return base_dsn + " " + pool_dsn

}
