package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	POSTGRES_CONN string `mapstructure:"POSTGRES_CONN"`
	PORT          string `mapstructure:"PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")

	if env == "production" {
		return EnvVars{
			POSTGRES_CONN: os.Getenv("POSTGRES_CONN"),
			PORT:          os.Getenv("PORT"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if config.POSTGRES_CONN == "" {
		err = errors.New("POSTGRES_CONN is required")
		return
	}

	if config.PORT == "" {
		err = errors.New("PORT is required")
		return
	}

	return
}
