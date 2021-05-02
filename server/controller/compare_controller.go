package controller

import (
	"encoding/json"
	"net/http"

	"github.com/serpedious/automatic-trading-system/server/compare/usecase"
)

func CompareTicker(w http.ResponseWriter, r *http.Request) {
	ticker_coincheck := usecase.GetAllTicker()
	js, _ := json.Marshal(ticker_coincheck)
	w.Write([]byte(js))
}
