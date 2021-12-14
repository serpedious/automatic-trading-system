package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/bitflyer/usecase"
	"github.com/serpedious/automatic-trading-system/server/config"
	"github.com/serpedious/automatic-trading-system/server/utils"
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

func GetBalanceHistory(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	params := map[string]string{
		"currency_code": "JPY",
		"count":         "10",
		"before":        "0",
		"after":         "0",
	}
	balance_history_data, _ := apiClient.GetBalanceHistory(params)
	js, _ := json.Marshal(balance_history_data)
	w.Write([]byte(js))
}

func Ticker(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	ticker_data, _ := apiClient.GetTicker(config.Config.ProductCode)
	js, _ := json.Marshal(ticker_data)
	w.Write([]byte(js))
}

func GetMyAssets(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	markets_list := []string{"BTC_JPY", "XRP_JPY", "ETH_JPY", "XLM_JPY", "MONA_JPY"}
	ticker_data := []*bitflyer.AssetsTicker{}
	for _, product_code := range markets_list {
		ticker, _ := apiClient.GetAssetsTicker(product_code)
		ticker_data = append(ticker_data, ticker)
	}
	balance_data, _ := apiClient.GetBalance()
	my_assets, _ := usecase.CalcMyAssets(ticker_data, balance_data)
	js, _ := json.Marshal(my_assets)
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
	fmt.Println("*******************************")
	fmt.Println(r.Body)
	apiClient := usecase.CreateClient()
	var order bitflyer.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	// order := &bitflyer.Order{
	// 	ProductCode:     "XRP_JPY",
	// 	ChildOrderType:  "MARKET",
	// 	Side:            "BUY",
	// 	Size:            size,
	// 	MinuteToExpires: 1,
	// 	TimeInForce:     "GTC",
	// }
	fmt.Println(order)
	fmt.Println("**********************")
	sendorder_data, _ := apiClient.SendOrder(&order)
	js, _ := json.Marshal(sendorder_data)
	w.Write([]byte(js))
}

var apiValidPath = regexp.MustCompile("^/api/candle/$")

func ApiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			utils.APIError(w, "Not found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func ApiCandleHandler(w http.ResponseWriter, r *http.Request) {
	productCode := r.URL.Query().Get("product_code")
	if productCode == "" {
		utils.APIError(w, "No product_code params", http.StatusBadRequest)
		return
	}
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 1000
	}

	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1m"
	}
	durationTime := config.Config.Durations[duration]

	df, _ := usecase.GetAllCandle(productCode, durationTime, limit)

	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
