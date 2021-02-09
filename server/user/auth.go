package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/serpedious/automatic-trading-system/server/utils"
)

func VerifyEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.ResponseByJSON(w, "authentication is ok!")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var errorObject utils.Error

		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		fmt.Println("bearerToken: ", bearerToken)

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("An error has occured")
				}
				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				utils.ErrorInResponse(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				utils.ErrorInResponse(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "This token is invalid"
			return
		}
	})
}
