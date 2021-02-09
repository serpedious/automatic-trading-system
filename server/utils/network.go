package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func ResponseByJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
	return
}

func ErrorInResponse(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
	return
}
