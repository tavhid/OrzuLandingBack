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
	category := r.URL.Query().Get("category")

	
	pageLimit, err := strconv.Atoi(r.URL.Query().Get("pageLimit"))
	if err != nil {
		pageLimit = 16
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	merchant, maxPage, err := h.merchant.GetList(search, city, category, uint(page), uint(pageLimit))

	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = map[string]interface{}{
		"merchants_list": merchant,
		"max_page":       (maxPage + int64(pageLimit) - 1) / int64(pageLimit),
	}
}
