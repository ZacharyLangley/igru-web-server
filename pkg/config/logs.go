package config

import "go.uber.org/zap"

type Logger struct {
	Development bool   `mapstructure:"development"`
	Level       string `mapstructure:"level"`
}

func (l Logger) Setup() (err error) {
	var cfg zap.Config
	if l.Development {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.Level, err = zap.ParseAtomicLevel(l.Level)
	if err != nil {
		return
	}
	var logger *zap.Logger
	logger, err = cfg.Build()
	if err != nil {
		return
	}
	logger.Info("Set up logger", zap.String("level", l.Level))
	zap.ReplaceGlobals(logger)
	return
}
