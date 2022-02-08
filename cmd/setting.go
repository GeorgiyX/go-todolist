package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Settings struct {
	MODE       string `env:"MODE" envDefault:"debug"`
	APIPort    uint   `env:"API_PORT" envDefault:"8080"`
	DBHost     string `env:"PG_HOST" envDefault:"localhost"`
	DBPort     uint   `env:"PG_PORT" envDefault:"5432"`
	DBName     string `env:"PG_DB_NAME" envDefault:"forum"`
	DBUser     string `env:"PG_FORUM_USER" envDefault:"forum_user"`
	DBPassword string `env:"PG_PASSWORD" envDefault:"forum_user_password"`
	DSN        string
	APIAddr    string
	APIPrefix  string
}

func LoadSettings() *Settings {
	settings := Settings{}

	if err := env.Parse(&settings); err != nil {
		fmt.Printf("%+v\n", err)
	}

	settings.DSN = fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		settings.DBHost, settings.DBPort, settings.DBUser, settings.DBPassword, settings.DBName)
	settings.APIAddr = fmt.Sprintf("0.0.0.0:%v", settings.APIPort)
	settings.APIPrefix = "/api"

	fmt.Printf("Server settings: %+v\n", settings)
	return &settings
}
