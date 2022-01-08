package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("healthy")
	ok := "healthy"
	js, _ := json.Marshal(ok)
	w.Write([]byte(js))
}
