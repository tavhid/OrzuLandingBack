package main

import (
	"cl/internal/service"
	"cl/internal/storage"
	"cl/internal/transport/http/handlers"
	"cl/internal/transport/http/router"
	"cl/pkg/bootstrap/http"
	"cl/pkg/cache"
	"cl/pkg/config"
	"cl/pkg/databases"
	"cl/pkg/gateways"
	"cl/pkg/logger"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		config.Module,
		logger.Module,
		databases.PostgresModule,
		databases.RedisModule,

		service.Module,
		storage.Module,
		handlers.Module,
		router.Module,
		gateways.Module,
		cache.Module,
	)

	app.Run()
}
