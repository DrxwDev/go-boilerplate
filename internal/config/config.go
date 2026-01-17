package config

import (
	"os"
	"time"
)

func LoadAppConfig() *AppConfig {
	return &AppConfig{
		PORT: os.Getenv("PORT"),
	}
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		URL: os.Getenv("DB_DSN"),
	}
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		ReadTimeout:       10 * time.Second,
		WriteTimout:       10 * time.Second,
		IdleTimeot:        90 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
}
