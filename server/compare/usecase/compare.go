package usecase

import (
	"encoding/json"

	coincheckgo "github.com/Akagi201/coincheckgo"

	"github.com/serpedious/automatic-trading-system/server/bitflyer/usecase"
	"github.com/serpedious/automatic-trading-system/server/compare"
	"github.com/serpedious/automatic-trading-system/server/config"
)

func GetAllTicker() *compare.AllTicker {
	apiClient := usecase.CreateClient()
	bitflyer, _ := apiClient.GetTicker(config.Config.ProductCode)

	coincheck := GetTickerFromCoinCheck()

	allticker := compare.AllTicker{Ticker: *bitflyer, CoinCheckTicker: *coincheck}
	return &allticker
}

func GetTickerFromCoinCheck() *compare.CoinCheckTicker {
	client := new(coincheckgo.CoinCheck).NewClient(config.Config.CoinCheckApiKey, config.Config.CoinCheckApiSecret)
	resp := client.Ticker.All()
	var tick_result compare.CoinCheckTicker
	json_err := json.Unmarshal([]byte(resp), &tick_result)
	if json_err != nil {
		panic(json_err)
	}
	return &tick_result
}
