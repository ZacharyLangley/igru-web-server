package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Database   int    `mapstructure:"database"`
	MaxRetries int    `mapstructure:"maxREtries"`
}

func (c Cache) Options() *redis.Options {
	return &redis.Options{
		Addr:       fmt.Sprintf("%s:%d", c.Host, c.Port),
		Username:   c.Username,
		Password:   c.Password,
		DB:         c.Database,
		MaxRetries: c.MaxRetries,
	}
}
