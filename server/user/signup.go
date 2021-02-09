package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/serpedious/automatic-trading-system/server/tool"
	"github.com/serpedious/automatic-trading-system/server/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user User
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

	sql_query := "INSERT INTO USERS(EMAIL, PASSWORD) VALUES($1, $2) RETURNING ID;"

	err = tool.Db.QueryRow(sql_query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		error.Message = "server error"
		utils.ErrorInResponse(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")

	utils.ResponseByJSON(w, user)
}
