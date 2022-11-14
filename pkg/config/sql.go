package config

import (
	"fmt"
	"strings"
)

type SQL struct {
	Host             string                 `mapstructure:"host"`
	Port             int                    `mapstructure:"port"`
	Database         string                 `mapstructure:"database"`
	Username         string                 `mapstructure:"username"`
	Password         string                 `mapstructure:"password"`
	MaxRetries       int                    `mapstructure:"maxRetries"`
	DisableMigration bool                   `mapstructure:"disableMigration"`
	Args             map[string]interface{} `mapstructure:",remain"`
}

func (d SQL) DSN() (string, error) {
	// Example DSN: "host=postgres port=5432 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"
	parts := make([]string, 0, 10)
	if d.Host != "" {
		parts = append(parts, fmt.Sprintf("host=%s", d.Host))
	}
	if d.Port > 0 {
		parts = append(parts, fmt.Sprintf("port=%d", d.Port))
	}
	if d.Database != "" {
		parts = append(parts, fmt.Sprintf("dbname=%s", d.Database))
	}
	if d.Username != "" {
		parts = append(parts, fmt.Sprintf("user=%s", d.Username))
	}
	if d.Password != "" {
		parts = append(parts, fmt.Sprintf("password=%s", d.Password))
	}
	for key, rawValue := range d.Args {
		switch x := rawValue.(type) {
		case string:
			parts = append(parts, fmt.Sprintf("%s=%s", key, x))
		case int:
			parts = append(parts, fmt.Sprintf("%s=%d", key, x))
		default:
			return "", fmt.Errorf("unknown args type for %s: %T", key, rawValue)
		}
	}
	return strings.Join(parts, " "), nil
}
