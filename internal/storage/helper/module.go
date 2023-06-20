package helper

import (
	"cl/internal/storage/helper/otp"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	otp.Module,
)
