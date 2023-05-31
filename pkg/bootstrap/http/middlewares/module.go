package middlewares

import (
	"cl/pkg/config"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(NewMiddleware)

// Middleware ...
type Middleware interface {
	JWT(next http.HandlerFunc) http.HandlerFunc
	CORS(next http.HandlerFunc) http.HandlerFunc
}

// Dependencies ...
type Dependencies struct {
	fx.In

	Config   *config.Config
	Logger   *logrus.Logger
	Postgres *gorm.DB
}

type provider struct {
	config   *config.Config
	logger   *logrus.Logger
	postgres *gorm.DB
}

// NewMiddleware ...
func NewMiddleware(params Dependencies) Middleware {
	return &provider{
		config:   params.Config,
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
