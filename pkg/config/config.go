package config

import (
	"flag"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "default.yml", "Path to your config file")
}

func New(conf any) error {
	if !flag.Parsed() {
		flag.Parse()
	}
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(conf); err != nil {
		return err
	}
	return nil
}
