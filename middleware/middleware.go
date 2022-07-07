package middleware

import (
	"encoding/json"
	"fmt"
	_const "jwt-practice/const"
	"jwt-practice/entity"
	"jwt-practice/service"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			var err entity.Error
			err = service.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(_const.SecretKey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err entity.Error
			err = service.SetError(err, "Your Token has been expired")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}
		var reserr entity.Error
		reserr = service.SetError(reserr, "Not Authorized")
		json.NewEncoder(w).Encode(err)
	}
}
