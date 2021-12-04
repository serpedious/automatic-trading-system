package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func ResponseByJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func ErrorInResponse(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, ErrorMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: ErrorMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}
