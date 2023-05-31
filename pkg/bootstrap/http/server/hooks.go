package server

import (
	"context"
	"net/http"

	"go.uber.org/fx"
)

// ModuleLifecycleHooks ...
var ModuleLifecycleHooks = fx.Invoke(RegisterHooks)

// RegisterHooks ...
func RegisterHooks(lifecycle fx.Lifecycle, server *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
