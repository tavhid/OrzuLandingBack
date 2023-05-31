package handlers

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
	"encoding/json"
	"net/http"
)

// HSendOTP ...
func (h *Handler) HSendApplication(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	var application structs.Application

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&application)

	if err != nil {
		err = response.ErrBadRequest
		resp.Message = err.Error()
		return
	}

	err = h.otp.SendOTP(application)
	if err != nil {

		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
}
