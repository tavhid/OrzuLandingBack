package http

import (
	"cl/pkg/bootstrap/http/middlewares"
	"cl/pkg/bootstrap/http/misc"
	"cl/pkg/bootstrap/http/server"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	middlewares.Module,
	misc.Module,

	server.ModuleLifecycleHooks,
	server.ServerModule,
)
