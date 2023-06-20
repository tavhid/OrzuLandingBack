package handlers

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
	"encoding/json"
	"net/http"
)

// HSendApplication ...
func (h *Handler) HSendApplication(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	var data map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		err = response.ErrBadRequest
		resp.Message = err.Error()
		return
	}

	if data["type_of_client"] == "merchant" {
		var newMerchant structs.MerchantApplication
		marshalData, _ := json.Marshal(data)
		err = json.Unmarshal(marshalData, &newMerchant)
		err = newMerchant.Validate()

	} else if data["type_of_client"] == "user" {
		var newUser structs.UserApplication
		marshalData, _ := json.Marshal(data)
		err = json.Unmarshal(marshalData, &newUser)
		err = newUser.Validate()
	} else {
		err = response.ErrBadRequest
		return
	}
	if err != nil {

		resp.Message = err.Error()
		return
	}

	err = h.otp.SendOTP(data["phone"].(string), "", data)
	if err != nil {

		resp.Message = err.Error()
		return
	}
	resp.Message = response.ErrSuccess.Error()
}
