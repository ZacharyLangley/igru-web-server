package config

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Logger struct {
	Development bool   `mapstructure:"development"`
	Level       string `mapstructure:"level"`
}

func SetupCobraLogger(cmd *cobra.Command, args []string) error {
	l := Logger{
		Development: viper.GetBool("dev-logs"),
		Level:       viper.GetString("level"),
	}
	// Get command name
	nameComponents := make([]string, 0, 1)
	currentCmd := cmd
	for {
		nameComponents = append([]string{currentCmd.Name()}, nameComponents...)
		currentCmd = currentCmd.Parent()
		if currentCmd == nil || currentCmd == currentCmd.Root() {
			break
		}
	}
	name := strings.Join(nameComponents, "-")

	logger, err := l.setup()
	if err != nil {
		return err
	}
	logger = logger.Named(name)
	logger.Info("Set up logger", zap.String("level", l.Level))
	zap.ReplaceGlobals(logger)
	return nil
}

func (l Logger) setup() (logger *zap.Logger, err error) {
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
	logger, err = cfg.Build()
	return
}
