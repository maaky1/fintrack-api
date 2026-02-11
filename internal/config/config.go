package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return nil, err
		}
	}

	v.SetDefault("APP_PORT", "8080")

	dsn := v.GetString("DATABASE_URL")
	if dsn == "" {
		dsn = v.GetString("database.url")
	}

	if dsn == "" {
		return nil, errors.New("DATABASE_URL (or database.url) is required")
	}

	return v, nil
}
