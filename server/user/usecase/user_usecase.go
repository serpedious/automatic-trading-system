package usecase

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/serpedious/automatic-trading-system/server/user"
)

func CreateToken(user user.User) (string, error) {
	_ = godotenv.Load()

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "__init__",
	})

	spew.Dump(token)
	tokenString, err := token.SignedString([]byte(secret))

	fmt.Println("--------------------")
	fmt.Println("tokenString: ", tokenString)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
