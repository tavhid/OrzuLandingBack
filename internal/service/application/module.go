package application

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewApplication)

// Application ...
type Application interface {
	Create(fullName, phone string, regionID, applicationTypeID uint) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Postgres *gorm.DB
	Logger   *logrus.Logger
}

type provider struct {
	postgres *gorm.DB
	logger   *logrus.Logger
}

// NewInfo ...
func NewApplication(params Dependencies) Application {
	return &provider{
		postgres: params.Postgres,
		logger:   params.Logger,
	}
}
