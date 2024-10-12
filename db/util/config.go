package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstruct:"DB_DRIVER"`
	DBSource            string        `mapstruct:"DB_SOURCE"`
	ServerAddress       string        `mapstruct:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstruct:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstruct:"ACCESS_TOKEN_DURATION"`
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
