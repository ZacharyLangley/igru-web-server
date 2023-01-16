package config

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Client struct {
	v *viper.Viper
}

func (c *Client) Get() (string, error) {
	sessionToken := c.v.GetString("token")
	if sessionToken == "" {
		return "", fs.ErrNotExist
	}
	return sessionToken, nil
}

func (c *Client) Set(sessionToken string) error {
	c.v.Set("token", sessionToken)
	c.v.Set("refreshed", time.Now())
	return c.v.WriteConfig()
}

func NewClient() (*Client, error) {
	v := viper.New()
	v.AutomaticEnv()
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	v.SetConfigName("igru")
	v.AddConfigPath(configDir)
	v.AddConfigPath(".")
	err = v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			v.SetConfigFile(path.Join(configDir, "igru.yaml"))
			return &Client{v: v}, nil
		}
		zap.L().Info("error reading config", zap.Error(err), zap.String("type", fmt.Sprintf("%T", err)))
		return nil, err
	}
	return &Client{v: v}, nil
}
