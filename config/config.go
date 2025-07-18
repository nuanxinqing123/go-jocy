package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"go-jocy/config/autoload"
)

type Configuration struct {
	App autoload.App `mapstructure:"app" json:"app" yaml:"app"`
}

var (
	GinConfig Configuration
	GinLOG    *zap.Logger
	GinVP     *viper.Viper
)
