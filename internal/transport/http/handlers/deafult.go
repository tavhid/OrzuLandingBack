package handlers

import (
	"cl/pkg/bootstrap/http/misc/response"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// HPingPong — test route to check server status
func (h *Handler) HPingPong(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = "Backend orzu landing pong!"
}

// HNotImplementation ...
func (h *Handler) HNotImplementation(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	resp.Message = response.ErrNotImplementation.Error()
}

// ServeSwaggerFiles ...
func (h *Handler) ServeSwaggerFiles(w http.ResponseWriter, r *http.Request) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if strings.Contains(r.URL.String(), "yaml") {
		http.ServeFile(w, r, pwd+"/../pkg/docs/swagger.yaml")
		return
	}

	http.ServeFile(w, r, pwd+"/../pkg/docs/swagger.json")
}
