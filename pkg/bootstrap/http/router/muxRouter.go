package router

import (
	"log"
	"net/http"
)

// ChainsMiddlewares ...
func (h *HTTPRouter) ChainsMiddlewares(method string, handler http.HandlerFunc, mw ...HTTPMw) http.HandlerFunc {
	mws := make([]HTTPMw, 0)
	mws = append(mws, mw...)

	for _, httpMw := range mws {
		handler = httpMw(handler)
	}

	return handler
}

// Handle ...
func (h *HTTPRouter) Handle(method string, pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	handler = h.ChainsMiddlewares(method, handler, mw...)
	h.router.HandleFunc(pattern, handler).Methods(method, "OPTIONS")
}

func (h *HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	h.router.ServeHTTP(w, r)
}

// GET ...
func (h *HTTPRouter) GET(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("GET", pattern, handler, mw...)
}

// POST ...
func (h *HTTPRouter) POST(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("POST", pattern, handler, mw...)
}

// PUT ...
func (h *HTTPRouter) PUT(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("PUT", pattern, handler, mw...)
}

// DELETE ...
func (h *HTTPRouter) DELETE(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("DELETE", pattern, handler, mw...)
}
