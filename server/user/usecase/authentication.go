package usecase

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticateToken(bearerToken []string) (interface{}, error) {
	if len(bearerToken) != 2 {
		return nil, fmt.Errorf("token length is inaccurate")
	}
	authToken := bearerToken[1]
	token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("an error has occured")
		}
		return []byte("secret"), nil
	})

	if !token.Valid {
		return nil, fmt.Errorf("token validation error")
	}

	return token, error
}
