package response

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewResponse)

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger *logrus.Logger
}

// Response ...
type Response struct {
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`

	contentType string `json:"-"`
	logger      *logrus.Logger
}

// NewResponse ...
func NewResponse(params Dependencies) *Response {
	return &Response{
		logger: params.Logger,
	}
}
