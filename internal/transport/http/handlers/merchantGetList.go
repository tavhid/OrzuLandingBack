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
	pageLimit, err := strconv.Atoi(r.URL.Query().Get("pageLimit"))
	if err != nil {
		pageLimit = 16
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	category, err := strconv.Atoi(r.URL.Query().Get("category"))

	if err != nil {
		category = 0
	}

	merchant, err := h.merchant.GetList(search, city, uint(category), uint(page), uint(pageLimit))

	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = merchant
}
