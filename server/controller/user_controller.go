package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/serpedious/automatic-trading-system/server/tool"
	"github.com/serpedious/automatic-trading-system/server/user"
	"github.com/serpedious/automatic-trading-system/server/utils"
	"golang.org/x/crypto/bcrypt"
)


func Signin(w http.ResponseWriter, r *http.Request) {
	var jwt user.JWT
	var user user.User
	var error utils.Error

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

	var Db *sql.DB
	err := godotenv.Load()

	if err != nil {
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", tool.HOST, tool.USER, tool.DATABASE, tool.PASSWORD)
		Db, _ = sql.Open("postgres", dbURI)
	} else {
		Db, _ = sql.Open("postgres", "host=postgres_db port=5432 user="+tool.USER+" password="+tool.PASSWORD+" dbname="+tool.DATABASE+" sslmode=disable")
	}

	row := Db.QueryRow("SELECT * FROM USERS WHERE email=$1;", user.Email)
	err = row.Scan(&user.ID, &user.Email, &user.Password)

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

	token, err := CreateToken(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token
	defer Db.Close()

	utils.ResponseByJSON(w, jwt)
}

func CreateToken(user user.User) (string, error) {
	var err error

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

func Signup(w http.ResponseWriter, r *http.Request) {
	var user user.User
	var error utils.Error

	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is required"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is required"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	fmt.Println(user)
	fmt.Println("------------------------")
	spew.Dump(user)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("password: ", user.Password)
	fmt.Println("hash password: ", hash)

	user.Password = string(hash)
	fmt.Println("converted password: ", user.Password)

	sql_query := "INSERT INTO USERS(EMAIL, PASSWORD) VALUES($1, $2) RETURNING id;"

	var Db *sql.DB

	err = godotenv.Load()

	if err != nil {
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", tool.HOST, tool.USER, tool.DATABASE, tool.PASSWORD)
		Db, _ = sql.Open("postgres", dbURI)
	} else {
		Db, _ = sql.Open("postgres", "host=postgres_db port=5432 user="+tool.USER+" password="+tool.PASSWORD+" dbname="+tool.DATABASE+" sslmode=disable")
	}

	err = Db.QueryRow(sql_query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		error.Message = "server error"
		utils.ErrorInResponse(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	defer Db.Close()

	utils.ResponseByJSON(w, user)
}

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
					return nil, fmt.Errorf("an error has occured")
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
