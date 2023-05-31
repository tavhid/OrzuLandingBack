package sms

import (
	"cl/pkg/config"

	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewGSmsModule)

// GSms ...
type GSms interface {
	Send(phone, text string) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger *logrus.Logger
	Config *config.Config
}

type provider struct {
	logger *logrus.Logger
	config *config.Config
}

// NewGSmsModule ...
func NewGSmsModule(params Dependencies) GSms {
	return &provider{
		logger: params.Logger,
		config: params.Config,
	}
}
