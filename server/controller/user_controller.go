package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/serpedious/automatic-trading-system/server/tool"
	"github.com/serpedious/automatic-trading-system/server/user"
	"github.com/serpedious/automatic-trading-system/server/user/usecase"
	"github.com/serpedious/automatic-trading-system/server/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var jwt user.JWT
	var u user.User
	var error utils.Error

	json.NewDecoder(r.Body).Decode(&u)

	v := u.Validate()
	if v != "" {
		error.Message = v
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	password := u.Password

	Db := tool.NewDb()
	defer Db.Close()

	row := Db.QueryRow("SELECT * FROM USERS WHERE email=$1;", u.Email)
	err := row.Scan(&u.ID, &u.Email, &u.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "no existed user"
			utils.ErrorInResponse(w, http.StatusBadRequest, error)
		} else {
			log.Fatal(err)
		}
	}

	hasedPassword := u.Password

	err = bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))

	if err != nil {
		error.Message = "This is invalid password"
		utils.ErrorInResponse(w, http.StatusUnauthorized, error)
		return
	}

	token, err := usecase.CreateToken(u)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token


	utils.ResponseByJSON(w, jwt)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var u user.User
	var error utils.Error

	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&u)

	v := u.Validate()
	if v != "" {
		error.Message = v
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}


	spew.Dump(u)

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		log.Fatal(err)
	}

	u.Password = string(hash)
	sql_query := "INSERT INTO USERS(EMAIL, PASSWORD) VALUES($1, $2) RETURNING id;"

	Db := tool.NewDb()
	defer Db.Close()

	err = Db.QueryRow(sql_query, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		error.Message = "server error"
		utils.ErrorInResponse(w, http.StatusInternalServerError, error)
		return
	}

	u.Password = ""
	w.Header().Set("Content-Type", "application/json")

	utils.ResponseByJSON(w, u)
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
