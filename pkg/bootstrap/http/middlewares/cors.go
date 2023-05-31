package middlewares

import (
	"cl/pkg/bootstrap/http/misc/response"
	"encoding/json"
	"net/http"
)

func (m *provider) CORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			data, _ := json.Marshal(response.Build(response.ErrNoContent))
			w.Write(data)
			return
		}

		next(w, r)
	})
}
