package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySQLDSN string `mapstructure:"MYSQL_DSN"`
}

func LoadConfig() *Config {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config
	viper.Unmarshal(&config)

	return &config
}
