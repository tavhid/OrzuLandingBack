package storage

import (
	"cl/internal/storage/application"
	"cl/internal/storage/helper"
	"cl/internal/storage/merchant"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	merchant.Module,
	helper.Module,
	application.Module,
)
