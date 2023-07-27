package service

import (
	"cl/internal/service/application"
	"cl/internal/service/cft"
	"cl/internal/service/merchant"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	application.Module,
	merchant.Module,
	cft.Module,
)
