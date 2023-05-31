package storage

import (
	"cl/internal/storage/merchant"
	"cl/internal/storage/otp"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	otp.Module,
	merchant.Module,
)
