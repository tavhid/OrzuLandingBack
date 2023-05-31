package handlers

import (
	"net/http"
	"strconv"

	"cl/pkg/bootstrap/http/misc/response"
)

// HMerchantGet ...
func (h *Handler) HMerchantGet(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		err = response.ErrBadRequest
		resp.Message = err.Error()
		return
	}

	merchant, err := h.merchant.Get(uint(id))

	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = merchant
}
