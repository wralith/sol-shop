package config

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Http HttpConfig `mapstructure:"http"`
}

type HttpConfig struct {
	Port string `mapstructure:"port"`
}

func NewConfig() (Config, error) {
	v := viper.New()
	var config Config

	switch os.Getenv("GO_ENV") {
	case "PRODUCTION":
		v.SetConfigName("config.production")
	default:
		v.SetConfigName("config.dev")
	}

	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	// i.e: HTTP_PORT environment variable will replace http.port from the yaml file
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return config, errors.Join(errors.New("error while reading config"), err)
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return config, errors.Join(errors.New("error while unmarshaling config"), err)
	}

	return config, nil
}
