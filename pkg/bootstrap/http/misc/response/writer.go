package response

import (
	"encoding/json"
	"net/http"
)

// WriterJSON — method is used to write a JSON response to http.ResponseWriter
func (r *Response) WriterJSON(w http.ResponseWriter) error {

	//if panic don't write response
	if rec := recover(); rec != nil {
		//recover() - is disposable need do reusable

		r.logger.Error(rec)
		r.Message = ErrSomethingWentWrong.Error()

		//stop responsible process
		return nil
	}

	// Trying to convert the data to json format
	body, err := json.Marshal(r)
	if err != nil {
		r.logger.Error(err)
		r.Message = ErrInternalServer.Error()

		return err
	}

	// Set Content Type of response
	w.Header().Set("Content-Type", "application/json")
	if r.contentType != "" {
		w.Header().Set("Content-Type", r.contentType)
	}

	// Setup headers for CORS preflight
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	code := errorCode[r.Message]
	if code == 0 {
		code = 500
	}

	// Set status code of response
	w.WriteHeader(code)
	w.Write(body)

	return nil
}
