package controller

import (
	"encoding/json"
	"net/http"

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
