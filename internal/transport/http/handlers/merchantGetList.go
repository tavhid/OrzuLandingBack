package handlers

import (
	"net/http"
	"strconv"

	"cl/pkg/bootstrap/http/misc/response"
)

// HMerchantGetList ...
func (h *Handler) HMerchantGetList(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	search := r.URL.Query().Get("search")
	city := r.URL.Query().Get("city")
	category, err := strconv.Atoi(r.URL.Query().Get("category"))

	if err != nil {
		category = 0
	}

	merchant, err := h.merchant.GetList(search, city, uint(category))

	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = merchant
}
