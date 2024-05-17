package config

import (
	"fmt"

	"github.com/ardanlabs/conf/v3"
)

type config struct {
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
		// https://pkg.go.dev/github.com/jackc/pgx/v5@v5.3.0/pgxpool#ParseConfig
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

func New() (*config, error) {
	var c config = config{}
	if _, err := conf.Parse("", &c); err != nil {
		return &config{}, err
	}

	return &c, nil
}

func (c *config) AppName() string {
	return c.App.Name
}

func (c *config) DbDsn() string {
	// https://pkg.go.dev/github.com/jackc/pgx/v5@v5.5.5/pgxpool#ParseConfig
	base_dsn := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	pool_dsn := "pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s pool_health_check_period=%s pool_max_conn_lifetime_jitter=%s"
	dsn := base_dsn + " " + pool_dsn

	return fmt.Sprintf(dsn, c.Db.Host, c.Db.Port, c.Db.User, c.Db.Password, c.Db.Name, c.Db.SSLMode, c.Db.Pool.MaxConn, c.Db.Pool.MinConn, c.Db.Pool.MaxConnLifetime, c.Db.Pool.MaxConnIdleTime, c.Db.Pool.HealthCheckPeriod, c.Db.Pool.MaxConnLifetimeJitter)
}
