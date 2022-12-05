package config

import (
	"github.com/spf13/viper"
)

var configPath string

func New(conf any) error {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigFile(viper.GetString("config"))
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(conf); err != nil {
		return err
	}
	return nil
}
