package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Pretty bool   `yaml:"pretty"`
}

func NewConfig() *Config {
	c := &Config{}
	c.initConfig()
	return c
}

func (c *Config) initConfig() {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("app")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	v.Unmarshal(c)
}
