package controllers

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	ok := "healthy"
	js, _ := json.Marshal(ok)
	w.Write([]byte(js))
}
