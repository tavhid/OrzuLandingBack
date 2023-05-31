package middlewares

import (
	"cl/pkg/bootstrap/http/misc/response"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func (p *provider) JWT(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		baererHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(baererHeader) < 2 {
			response.Build(response.ErrAccessDenied).WriterJSON(rw)
			return
		}

		baererToken := baererHeader[1]

		token, err := jwt.Parse(baererToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			if err.Error() == "Token is expired" {
				response.Build(response.ErrAccessDenied).WriterJSON(rw)
				return
			}
			response.Build(response.ErrAccessDenied).WriterJSON(rw)
			return
		}

		tokenMap := token.Claims.(jwt.MapClaims)
		userID := int(tokenMap["user_id"].(float64))

		r.Header.Set("user_id", fmt.Sprint(userID))
		handler.ServeHTTP(rw, r)
	})
}
