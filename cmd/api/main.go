package main

import (
	"github.com/DrxwDev/boilerplate/internal/config"
	"github.com/DrxwDev/boilerplate/internal/database"
	"github.com/DrxwDev/boilerplate/internal/health"
	"github.com/DrxwDev/boilerplate/internal/server"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	app := fx.New(
		config.Module,
		database.Module,
		server.Module,
		health.Module,
	)

	app.Run()
}
