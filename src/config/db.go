package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DSN string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
			viper.GetString("DB_HOST"),
			viper.GetString("DB_USERNAME"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_NAME"),
			viper.GetInt("DB_PORT")),
	}
}
