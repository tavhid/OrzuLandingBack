package cft

import (
	"cl/internal/structs"
	"database/sql"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(Cft)

// CFT ...
type CFT interface {
	GetMerchantList() (merchantList []structs.Merchant, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Cft      *sql.DB
	Logger   *logrus.Logger
	Postgres *gorm.DB
}

type provider struct {
	cft      *sql.DB
	logger   *logrus.Logger
	postgres *gorm.DB
}

// Cft ...
func Cft(params Dependencies) CFT {
	return &provider{
		cft:      params.Cft,
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
