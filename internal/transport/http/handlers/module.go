package handlers

import (
	"cl/internal/storage/application"
	"cl/internal/storage/helper/otp"
	"cl/internal/storage/merchant"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger      *logrus.Logger
	Otp         otp.Otp
	Merchant    merchant.Merchant
	Application application.Application
}

// Handler ...
type Handler struct {
	logger      *logrus.Logger
	otp         otp.Otp
	merchant    merchant.Merchant
	application application.Application
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger:      params.Logger,
		otp:         params.Otp,
		merchant:    params.Merchant,
		application: params.Application,
	}
}
