package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	ApiAddress string `mapstructure:"API_ADDRESS"`
	DBSource   string `mapstructure:"DB_SOURCE"`
}

func NewConfig(log *zap.Logger) *Config {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.SetDefault("API_ADDRESS", ":8080")

	viper.ReadInConfig()
	viper.BindEnv("API_ADDRESS")
	viper.BindEnv("DB_SOURCE")

	var config Config
	err := viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("Cannot unmarshal config:", zap.Error(err))
	}

	return &config
}
