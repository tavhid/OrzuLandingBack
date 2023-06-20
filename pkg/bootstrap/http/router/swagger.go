package router

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectSwagger ...
func (h *HTTPRouter) ConnectSwagger(file func(w http.ResponseWriter, r *http.Request)) {
	swaggerOptions := middleware.SwaggerUIOpts{
		BasePath: "/landing",
		SpecURL:  "/landing/swagger.yaml",
	}
	sh := middleware.SwaggerUI(swaggerOptions, nil)

	h.GET("/swagger.yaml", file)
	h.GET("/swagger.json", file)

	h.router.Path("/docs").Handler(sh)
}
