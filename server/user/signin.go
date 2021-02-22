package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/serpedious/automatic-trading-system/server/tool"
	"github.com/serpedious/automatic-trading-system/server/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	HOST     = os.Getenv("DATABASE_HOST")
	// 	DATABASE = os.Getenv("POSTGRES_DB")
	// 	USER     = os.Getenv("POSTGRES_USER")
	// 	PASSWORD = os.Getenv("PGPASSWORD")
	// )

	var user User
	var error utils.Error
	var jwt JWT

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is required"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is required"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
	}
	password := user.Password
	fmt.Println("password: ", password)
	// var Db *sql.DB
	// Db, _ = sql.Open("postgres", "host="+HOST+" port=5432 user="+USER+" password="+PASSWORD+" dbname="+DATABASE+" sslmode=disable")

	tool.Connect()
	row := tool.Db.QueryRow("SELECT * FROM USERS WHERE email=$1;", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "no existed user"
			utils.ErrorInResponse(w, http.StatusBadRequest, error)
		} else {
			log.Fatal(err)
		}
	}

	hasedPassword := user.Password
	fmt.Println("hasedPassword: ", hasedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))

	if err != nil {
		error.Message = "This is invalid password"
		utils.ErrorInResponse(w, http.StatusUnauthorized, error)
		return
	}

	token, err := createToken(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	utils.ResponseByJSON(w, jwt)
}

func createToken(user User) (string, error) {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

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
