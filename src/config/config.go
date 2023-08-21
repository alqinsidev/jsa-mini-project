package config

import "github.com/spf13/viper"

type Config struct {
	DB DBConfig
}

func EnvConfig() *Config {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	return &Config{
		DB: LoadDBConfig(),
	}
}
