package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/markcheno/go-talib"
	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/bitflyer/usecase"
	"github.com/serpedious/automatic-trading-system/server/config"
	"github.com/serpedious/automatic-trading-system/server/utils"
)

func CleanUpStreamData() {
	for {
		time.Sleep(time.Hour * 12)
		err := usecase.CronDelete()
		if err != nil {
			log.Fatalln("clean up error")
		}
	}
}

func StreamIngectionData() {
	var tickerChannl = make(chan usecase.Ticker)
	apiClient := usecase.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	for ticker := range tickerChannl {
		log.Printf("action=StreamIngectionData, %v", ticker)
		for _, duration := range config.Config.Durations {
			isCreated := usecase.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated && duration == config.Config.TradeDuration {
				usecase.AutomaticNotification()
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

func GetAlert(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	alert_data, _ := apiClient.GetAlert()
	js, _ := json.Marshal(alert_data)
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

func CalcProfit(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	markets_list := []string{"BTC_JPY", "XRP_JPY", "ETH_JPY", "XLM_JPY", "MONA_JPY"}
	ticker_data := []*bitflyer.AssetsTicker{}
	for _, product_code := range markets_list {
		ticker, _ := apiClient.GetAssetsTicker(product_code)
		ticker_data = append(ticker_data, ticker)
	}
	balance_data, _ := apiClient.GetBalance()
	my_assets, _ := usecase.CalcMyAssets(ticker_data, balance_data)
	var all_myassets float64
	for i := 0; i < len(my_assets); i++ {
		all_myassets += float64(my_assets[i].Value)
	}
	jpy_value, _ := apiClient.GetBalance()
	jpy := jpy_value[0].Available
	profit_data, _ := apiClient.CalcProfit(all_myassets, jpy)
	data := int(profit_data)
	js, _ := json.Marshal(data)
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
	var order bitflyer.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	fmt.Println(order)
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
	// productCode := r.URL.Query().Get("product_code")
	// if productCode == "" {
	// 	productCode = config.Config.ProductCode
	// }
	// // strLimit := r.URL.Query().Get("limit")
	// limit, err := strconv.Atoi(strLimit)
	// if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
	// 	limit = 100
	// }

	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1m"
	}
	productCode := config.Config.ProductCode
	limit := 100
	durationTime := config.Config.Durations[duration]

	df, _ := usecase.GetAllCandle(productCode, durationTime, limit)

	firstTime := df.Candles[0].Time
	df.AddEvents(firstTime)
	df.AddRsi(14)

	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(js)
}

func GetRsi(w http.ResponseWriter, r *http.Request) {
	period := 14
	df, _ := usecase.GetAllCandle("BTC_JPY", time.Minute, 365)
	if period < 14 {
		fmt.Println("less data set compare with config")
	}
	lenCandles := len(df.Candles)
	if lenCandles <= period {
		return
	}
	var rsi_data []float64
	values := talib.Rsi(df.Closes(), period)
	lastValue := values[lenCandles-1]
	secondLastValue := values[lenCandles-2]
	rsi_data = append(rsi_data, lastValue, secondLastValue)
	js, _ := json.Marshal(rsi_data)
	w.Write([]byte(js))
}

func GetSignal(w http.ResponseWriter, r *http.Request) {
	limit := 1
	signal_data := usecase.GetSignalEventsByCount(limit)
	js, _ := json.Marshal(signal_data)
	w.Write([]byte(js))
}

func GetSignalAll(w http.ResponseWriter, r *http.Request) {
	limit := 100
	signal_data := usecase.GetSignalEventsByCount(limit)
	js, _ := json.Marshal(signal_data)
	w.Write([]byte(js))
}
