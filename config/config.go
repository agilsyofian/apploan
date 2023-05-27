package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppName                    string        `mapstructure:"APPNAME"`
	Environment                string        `mapstructure:"ENVIRONMENT"`
	DBDriver                   string        `mapstructure:"DB_DRIVER"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	MigrationURL               string        `mapstructure:"MIGRATION_URL"`
	ServerAddress              string        `mapstructure:"SERVER_ADDRESS"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AccessTokenRefreshDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	Assets                     string        `mapstructure:"ASSETS"`
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
