package config

import (
	"fmt"
	"strings"
)

type Database struct {
	Host          string                 `mapstructure:"host"`
	Port          int                    `mapstructure:"port"`
	Database      string                 `mapstructure:"database"`
	Username      string                 `mapstructure:"username"`
	Password      string                 `mapstructure:"password"`
	MaxRetries    int                    `mapstructure:"maxRetries"`
	MigrationPath string                 `mapstructure:"migrationPath"`
	Args          map[string]interface{} `mapstructure:",remain"`
}

func (d Database) SecureDSN() string {
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
		parts = append(parts, "password=XXXXXXXX")
	}
	for key, rawValue := range d.Args {
		switch x := rawValue.(type) {
		case string:
			parts = append(parts, fmt.Sprintf("%s=%s", key, x))
		case int:
			parts = append(parts, fmt.Sprintf("%s=%d", key, x))
		default:
			return ""
		}
	}
	return strings.Join(parts, " ")
}

func (d Database) DSN() (string, error) {
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
			return "", unknownArgsTypesError{
				Key:      key,
				RawValue: rawValue,
			}
		}
	}
	return strings.Join(parts, " "), nil
}

type unknownArgsTypesError struct {
	Key      string
	RawValue any
}

func (u unknownArgsTypesError) Error() string {
	return fmt.Sprintf("unknown args type for %s: %T", u.Key, u.RawValue)
}
