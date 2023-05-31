package misc

import (
	"cl/pkg/bootstrap/http/misc/response"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	response.Module,
)
