package handlers

import (
	"cl/internal/storage/merchant"
	"cl/internal/storage/otp"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger   *logrus.Logger
	Otp      otp.Otp
	Merchant merchant.Merchant
}

// Handler ...
type Handler struct {
	logger   *logrus.Logger
	otp      otp.Otp
	merchant merchant.Merchant
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger:   params.Logger,
		otp:      params.Otp,
		merchant: params.Merchant,
	}
}
