package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/serpedious/automatic-trading-system/server/tool"
	"github.com/serpedious/automatic-trading-system/server/user/usecase"
	"github.com/serpedious/automatic-trading-system/server/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := tool.GetUserIdFromCookie(w, r)
	fmt.Println("***********************")
	fmt.Println(userId)
	fmt.Println("***********************")
	var u usecase.User
	v, err := u.GetUser(userId)
	var error utils.Error
	if err != nil {
		error.Message = "failed to get user"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.ResponseByJSON(w, v)
}

func EditPass(w http.ResponseWriter, r *http.Request) {
	var u usecase.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return
	}
	defer r.Body.Close()
	userId := tool.GetUserIdFromCookie(w, r)

	err = u.EditPass(userId)
	var error utils.Error
	if err != nil {
		error.Message = "failed to update password"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}
	v := "completed update"

	w.WriteHeader(http.StatusOK)
	utils.ResponseByJSON(w, v)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var u usecase.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return
	}
	defer r.Body.Close()

	m := u.Validate()
	var error utils.Error
	if m != "" {
		error.Message = m
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	jwt, err := u.SignIn()
	if err != nil {
		error.Message = "no existed user or invalid password"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	atCookie := &http.Cookie{
		Name:     "access_token",
		Path:     "/",
		Value:    jwt,
		Secure:   true,
		SameSite: 4,
		Expires:  time.Now().Add(time.Minute * 15),
	}
	http.SetCookie(w, atCookie)
	w.WriteHeader(http.StatusOK)
	utils.ResponseByJSON(w, jwt)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var u usecase.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return
	}
	defer r.Body.Close()

	v := u.Validate()
	var error utils.Error
	if v != "" {
		error.Message = v
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	jwt, err := u.SignUp()
	if err != nil {
		error.Message = "faied to create user"
		utils.ErrorInResponse(w, http.StatusInternalServerError, error)
		return
	}

	atCookie := &http.Cookie{
		Name:     "access_token",
		Path:     "/",
		Value:    jwt,
		Secure:   true,
		SameSite: 4,
		Expires:  time.Now().Add(time.Minute * 15),
	}

	useremail := u.Email
	utils.Mail(useremail)
	http.SetCookie(w, atCookie)
	w.Header().Set("Content-Type", "application/json")
	utils.ResponseByJSON(w, u)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("access_token")
	c.MaxAge = -1
	http.SetCookie(w, c)

	w.WriteHeader(http.StatusOK)
}

func VerifyEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.ResponseByJSON(w, "authentication is ok!")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		var err utils.Error
		_, error := usecase.AuthenticateToken(bearerToken)
		if error != nil {
			err.Message = error.Error()
			utils.ErrorInResponse(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
