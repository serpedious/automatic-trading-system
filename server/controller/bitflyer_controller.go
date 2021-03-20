package controller

import (
	"encoding/json"
	"net/http"

	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/config"
)

func Balance(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	balance_data, _ := apiClient.GetBalance()
	js, _ := json.Marshal(balance_data)
	w.Write([]byte(js))
}

func Ticker(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker_data, _ := apiClient.GetTicker(config.Config.ProductCode)
	js, _ := json.Marshal(ticker_data)
	w.Write([]byte(js))
}

func GetExecution(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	execution_data, _ := apiClient.GetExecution(config.Config.ProductCode)
	js, _ := json.Marshal(execution_data)
	w.Write([]byte(js))
}

func Listorder(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
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
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
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
