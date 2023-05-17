package config

import (
	"fmt"
	"strings"

	"github.com/abelz123456/celestial-api/package/log"
	"github.com/spf13/viper"
)

func Init(path string) Config {
	logger := log.NewLog()
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigName("base")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigName("test")
		logger.PanicOnError(viper.ReadInConfig(), "", nil)
	}

	logger.PanicOnError(viper.Unmarshal(&cfg), "", nil)
	cfg.BasePath = path

	optimizeConfig(&cfg)
	return cfg
}

func optimizeConfig(cfg *Config) {
	if cfg.AppEnv == "" {
		cfg.AppEnv = "development"
	}

	if cfg.DevelopmentPort == "" {
		cfg.DevelopmentPort = "3000"
	}

	if !strings.HasPrefix(cfg.DevelopmentPort, ":") {
		cfg.DevelopmentPort = fmt.Sprintf(":%s", cfg.DevelopmentPort)
	}
}
