package handlers

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
	"net/http"
)

// HCheckOTP ...
func (h *Handler) HCheckOTP(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	code := r.URL.Query().Get("code")
	phone := r.URL.Query().Get("phone")

	data := structs.OtpInput{
		PhoneNumber: phone,
		OtpCode:     code,
	}

	err := h.otp.CheckOTP(data)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
}
