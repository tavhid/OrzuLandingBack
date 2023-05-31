package handlers

import (
	"net/http"

	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
)

// HGetInfo ...
func (h *Handler) HGetInfo(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	var returnData structs.MainData

	returnData.Limit = 50000

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = map[string]interface{}{
		"limit": returnData.Limit,
	}
}
