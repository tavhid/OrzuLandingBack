package handlers

import (
	"net/http"

	"cl/pkg/bootstrap/http/misc/response"
)

// HGetInfo ...
func (h *Handler) HGetInfo(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = map[string]interface{}{
		"limit": 50000,
	}
}