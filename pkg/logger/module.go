package logger

import "go.uber.org/fx"

// Module ...
var Module = fx.Provide(InitLogger)
