package config

import (
	"fmt"
	"github.com/ardanlabs/conf/v3"
)

type Config struct {
	App struct {
		Name string `conf:"env:APP_NAME,default:go-poc-example"`
	}
	Db struct {
		Url      string `conf:"env:DB_URL"`
		Name     string `conf:"env:DB_NAME,default:postgres"`
		Host     string `conf:"env:DB_HOST,default:localhost"`
		Port     int    `conf:"env:DB_PORT,default:5432"`
		User     string `conf:"env:DB_USER,default:postgres"`
		Password string `conf:"env:DB_PASSWORD,default:postgres"`
		SSLMode  string `conf:"env:DB_SSLMODE,default:disable"`
		// Valid log levels: https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5/tracelog#LogLevelFromString
		LogLevel string `conf:"env:DB_LOG_LEVEL,default:info"`
		/*
			https://pkg.go.dev/github.com/jackc/pgx#ConnConfig
			https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5/pgxpool#Config
		*/
		Pool struct {
			MaxConn               int    `conf:"env:DB_POOL_MAX_CONN,default:5"`
			MinConn               int    `conf:"env:DB_POOL_MIN_CONN,default:0"`
			MaxConnLifetime       string `conf:"env:DB_POOL_MAX_CONN_LIFETIME,default:5m"`
			MaxConnIdleTime       string `conf:"env:DB_POOL_MAX_CONN_IDLE_TIME,default:5m"`
			HealthCheckPeriod     string `conf:"env:DB_POOL_HEALTH_CHECK_PERIOD,default:5m"`
			MaxConnLifetimeJitter string `conf:"env:DB_POOL_MAX_CONN_LIFETIME_JITTER,default:30s"`
		}
	}
}

func NewConfig() (*Config, error) {
	var c Config = Config{}
	if _, err := conf.Parse("", &c); err != nil {
		return &Config{}, fmt.Errorf("fail to parse config: %w", err)
	}

	return &c, nil
}

func (c *Config) AppName() string {
	return c.App.Name
}

func (c *Config) DbDsn() string {
	/*
		https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
		https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5/pgxpool#ParseConfig
	*/

	baseDsn := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	poolDsn := "pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s pool_health_check_period=%s pool_max_conn_lifetime_jitter=%s"
	dsn := baseDsn + " " + poolDsn

	return fmt.Sprintf(dsn, c.Db.Host, c.Db.Port, c.Db.User, c.Db.Password, c.Db.Name, c.Db.SSLMode, c.Db.Pool.MaxConn, c.Db.Pool.MinConn, c.Db.Pool.MaxConnLifetime, c.Db.Pool.MaxConnIdleTime, c.Db.Pool.HealthCheckPeriod, c.Db.Pool.MaxConnLifetimeJitter)
}
