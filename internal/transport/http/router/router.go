package router

import (
	"cl/internal/transport/http/handlers"
	"cl/pkg/bootstrap/http/middlewares"
	transportHTTP "cl/pkg/bootstrap/http/router"
)

// NewRouter ..
func NewRouter(h *handlers.Handler, mw middlewares.Middleware) (router *transportHTTP.HTTPRouter) {
	router = transportHTTP.NewRouter()
	router.ConnectSwagger(h.ServeSwaggerFiles)

	router.GET("/ping", h.HPingPong, mw.CORS)
	// router.GET("/getInfo", h.HGetInfo, mw.CORS)
	router.POST("/sendApplication", h.HSendApplication, mw.CORS)
	router.GET("/checkOtp", h.HCheckOTP, mw.CORS)
	router.GET("/getMerchantList", h.HMerchantGetList, mw.CORS)
	router.GET("/getMerchant", h.HMerchantGet, mw.CORS)
	router.GET("/qr", h.HRedirectQR, mw.CORS)

	return
}
