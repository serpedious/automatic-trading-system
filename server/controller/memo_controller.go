package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/serpedious/automatic-trading-system/server/memo/usecase"
	"github.com/serpedious/automatic-trading-system/server/utils"
)

func CreateMemo(w http.ResponseWriter, r *http.Request) {
	var m usecase.Memo
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return
	}
	defer r.Body.Close()

	v := m.Validate()
	var error utils.Error
	if v != "" {
		error.Message = v
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	c := m.CreateMemo()
	if c != nil {
		error.Message = "failed create memo"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.ResponseByJSON(w, "created")
}

func DoneMemo(w http.ResponseWriter, r *http.Request) {
	var m usecase.DoneMemo
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	d := m.DoneMemo()
	var error utils.Error
	if d != nil {
		error.Message = "failed to switch done"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.ResponseByJSON(w, "switch status!")
}

func DeleteMemo(w http.ResponseWriter, r *http.Request) {
	var m usecase.DeleteMemo
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	d := m.DeleteMemo()
	var error utils.Error
	if d != nil {
		error.Message = "failed to delete memo"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.ResponseByJSON(w, "delete memo!")
}

func findUserId(claims map[string]interface{}) interface{} {
	for key, val := range claims {
		if key == "user_id" {
			fmt.Printf("Key: %v, value: %v\n", key, val)
			return val
		}
	}
	return nil
}

func GetAllMemos(w http.ResponseWriter, r *http.Request) {
	// TODO(refactoring)
	c, _ := r.Cookie("access_token")
	tknStr := c.Value
	claims := jwt.MapClaims{}
	secret := os.Getenv("JWT_SECRET")
	token, _ := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	fmt.Println(token)
	Id := findUserId(claims)
	userIdd := Id.(float64)
	userId := int(userIdd)


	var m usecase.GetMemo
	g, err := m.GetAllMemo(userId)
	var error utils.Error
	if err != nil {
		error.Message = "failed to get memos"
		utils.ErrorInResponse(w, http.StatusBadRequest, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.ResponseByJSON(w, g)
}
