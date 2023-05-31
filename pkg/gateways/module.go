package gateways

import (
	"cl/pkg/gateways/sms"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	sms.Module,
)
