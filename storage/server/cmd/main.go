package main

import (
	"github.com/ArkamFahry/hyperdrift/storage/server/api"
	"github.com/ArkamFahry/hyperdrift/storage/server/config"
	"github.com/ArkamFahry/hyperdrift/storage/server/database/migrations"
	"github.com/ArkamFahry/hyperdrift/storage/server/logger"
)

func main() {
	appConfig := config.NewConfig()

	appLogger := logger.NewLogger(appConfig)

	migrations.NewMigrations(appConfig, appLogger)

	api.NewApi(appLogger, appConfig)
}
