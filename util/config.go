package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER       string        `mapstructure:"DB_DRIVER"`
	DB_SERVICE      string        `mapstructure:"DB_SERVICE"`
	SERVICE_ADDRESS string        `mapstructure:"SERVICE_ADDRESS"`
	SECRET          string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	ACCESS_DURATION time.Duration `mapstructure:"TOKEN_ACCESS_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
