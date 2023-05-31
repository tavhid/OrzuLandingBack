package cache

import (
	"cl/pkg/cache/otp"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	otp.Module,
)
