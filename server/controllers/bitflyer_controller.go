package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/bitflyer/usecase"
	"github.com/serpedious/automatic-trading-system/server/config"
)

func StreamIngectionData() {
	var tickerChannl = make(chan usecase.Ticker)
	apiClient := usecase.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	for ticker := range tickerChannl {
		log.Printf("action=StreamIngectionData, %v", ticker)
		for _, duration := range config.Config.Durations {
			isCreated := usecase.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated && duration == config.Config.TradeDuration {
				fmt.Println(isCreated)
			}
		}
	}
}

func GetRealTimeTicker() {
	apiClient := usecase.New(config.Config.ApiKey, config.Config.ApiSecret)
	tickerChannel := make(chan usecase.Ticker)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker)
	}
}

func Balance(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	balance_data, _ := apiClient.GetBalance()
	js, _ := json.Marshal(balance_data)
	w.Write([]byte(js))
}

func Ticker(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	ticker_data, _ := apiClient.GetTicker(config.Config.ProductCode)
	js, _ := json.Marshal(ticker_data)
	w.Write([]byte(js))
}

func GetExecution(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	execution_data, _ := apiClient.GetExecution(config.Config.ProductCode)
	js, _ := json.Marshal(execution_data)
	w.Write([]byte(js))
}

func Listorder(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	i := "JRF2hogehoge"
	params := map[string]string{
		"product_code":              config.Config.ProductCode,
		"child_order_acceptance_id": i,
	}
	listorder_data, _ := apiClient.ListOrder(params)
	js, _ := json.Marshal(listorder_data)
	w.Write([]byte(js))
}

func SendOrder(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "LIMIT",
		Side:            "BUY",
		Price:           6000000,
		Size:            0.001,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
	sendorder_data, _ := apiClient.SendOrder(order)
	js, _ := json.Marshal(sendorder_data)
	w.Write([]byte(js))
}