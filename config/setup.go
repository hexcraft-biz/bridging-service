package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/hexcraft-biz/env"
	envRedis "github.com/hexcraft-biz/env/redis"
	"github.com/jmoiron/sqlx"
)

//================================================================
// Env
//================================================================
type Env struct {
	*env.Prototype
}

func FetchEnv() (*Env, error) {
	if e, err := env.Fetch(); err != nil {
		return nil, err
	} else {
		return &Env{
			Prototype: e,
		}, nil
	}
}

//================================================================
// Config
//================================================================
type Config struct {
	*Env
	DB    *sqlx.DB
	Redis *redis.Client
}

func Load() (*Config, error) {
	e, err := FetchEnv()
	if err != nil {
		return nil, err
	}

	return &Config{Env: e}, nil
}

func (cfg *Config) InitRedis() error {
	var err error

	cfg.Redis, err = envRedis.NewRedisClient()
	return err
}

func (cfg *Config) DBOpen(init bool) error {
	var err error

	cfg.DBClose()
	cfg.DB, err = cfg.MysqlConnectWithMode(init)

	return err
}

func (cfg *Config) DBClose() {
	if cfg.DB != nil {
		cfg.DB.Close()
	}
}
