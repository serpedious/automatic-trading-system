package usecase

import (
	"encoding/json"

	coincheckgo "github.com/Akagi201/coincheckgo"

	"github.com/serpedious/automatic-trading-system/server/compare"
	"github.com/serpedious/automatic-trading-system/server/config"
)

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
