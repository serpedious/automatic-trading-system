package controller

import (
	"encoding/json"
	"net/http"

	"github.com/serpedious/automatic-trading-system/server/compare/usecase"
)

func CompareTicker(w http.ResponseWriter, r *http.Request) {
	ticker_data := usecase.GetTickerFromCoinCheck()
	js, _ := json.Marshal(ticker_data)
	w.Write([]byte(js))
}
