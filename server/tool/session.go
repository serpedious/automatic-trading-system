package tool

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func findUserId(claims map[string]interface{}) interface{} {
	for key, val := range claims {
		if key == "user_id" {
			fmt.Printf("Key: %v, value: %v\n", key, val)
			return val
		}
	}
	return nil
}

func GetUserIdFromCookie(w http.ResponseWriter, r *http.Request) int {
	c, _ := r.Cookie("access_token")
	tknStr := c.Value
	claims := jwt.MapClaims{}
	secret := os.Getenv("JWT_SECRET")
	_, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	Id := findUserId(claims)
	userIdd := Id.(float64)
	userId := int(userIdd)

	return userId
}
