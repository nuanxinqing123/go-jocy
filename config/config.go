package config

import (
	"github.com/bluele/gcache"
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
	GinCache  gcache.Cache // 变更为Redis存储
	GinVP     *viper.Viper
)
