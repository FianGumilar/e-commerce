package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Env string
	}
	Server struct {
		Host string
		Port string
	}
	Postgres struct {
		Host string
		Port string
		User string
		Pass string
		Name string
		Ssl  string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file: ", err)
	}

	if appConfig == nil {
		appConfig = &AppConfig{}
		initApp(appConfig)
		initPostgres(appConfig)
	}
	return appConfig
}
