package usecase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/markcheno/go-talib"

	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/config"
)

type Ticker bitflyer.Ticker

type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func New(key, secret string) *APIClient {
	apiClient := &APIClient{key, secret, &http.Client{}}
	return apiClient
}

func CreateClient() *APIClient {
	apiClient := New(config.Config.ApiKey, config.Config.ApiSecret)
	return apiClient
}

func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(bitflyer.BaseURL)
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (api *APIClient) GetBalance() ([]bitflyer.Balance, error) {
	url := "me/getbalance"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance []bitflyer.Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance, nil
}

func subOnlyCrptoName(product_code string) string {
	slice := strings.Split(product_code, "_")
	return slice[0]
}

func CalcMyAssets(ticker_data []*bitflyer.AssetsTicker, balance_data []bitflyer.Balance) ([]bitflyer.MyAssets, error) {
	var myasset bitflyer.MyAssets
	var myassets []bitflyer.MyAssets

	set_name_ltp := map[string]float64{}
	for i := 0; i < len(ticker_data); i++ {
		v := ticker_data[i]
		symbol := subOnlyCrptoName(v.ProductCode)
		set_name_ltp[symbol] = v.Ltp
	}
	for i := 0; i < len(balance_data); i++ {
		curr_balance := balance_data[i]
		for key, value := range set_name_ltp {
			fmt.Println(key)
			if key == curr_balance.CurrentCode {
				jpy_value := curr_balance.Available * value
				myasset.Crpto = key
				myasset.Amount = curr_balance.Available
				myasset.Price = value
				myasset.Value = int(jpy_value)
				myassets = append(myassets, myasset)
			}
		}
	}

	return myassets, nil
}

func (api *APIClient) GetBalanceHistory(query map[string]string) ([]bitflyer.BalanceHistory, error) {
	url := "me/getbalancehistory"
	resp, err := api.doRequest("GET", url, query, nil)
	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance_history []bitflyer.BalanceHistory
	err = json.Unmarshal(resp, &balance_history)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance_history, nil
}

func (t *Ticker) GetMidPrice() float64 {
	return (t.BestBid + t.BestAsk) / 2
}

func (t *Ticker) DateTime() time.Time {
	dateTime, err := time.Parse(time.RFC3339, t.Timestamp)
	if err != nil {
		log.Printf("action=DateTime, err=%s", err.Error())
	}
	return dateTime
}

func (t *Ticker) TruncateDateTime(duration time.Duration) time.Time {
	return t.DateTime().Truncate(duration)
}

func (api *APIClient) GetTicker(productCode string) (*bitflyer.Ticker, error) {
	url := "ticker"
	resp, err := api.doRequest("GET", url, map[string]string{"product_code": productCode}, nil)
	if err != nil {
		return nil, err
	}
	var ticker bitflyer.Ticker
	err = json.Unmarshal(resp, &ticker)
	if err != nil {
		return nil, err
	}
	return &ticker, nil
}

func getWithDrawals(resp []byte, alert *bitflyer.Alert) {
	err := json.Unmarshal(resp, &alert.With)
	if err != nil {
		fmt.Println(err)
	}
}

func getDeposit(resp []byte, alert *bitflyer.Alert) {
	err := json.Unmarshal(resp, &alert.Dep)
	if err != nil {
		fmt.Println(err)
	}
}

func (api *APIClient) GetAlert() (*bitflyer.Alert, error) {
	url := "me/getwithdrawals"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return nil, err
	}
	alert := &bitflyer.Alert{}
	getWithDrawals(resp, alert)

	url = "me/getdeposits"
	deposit_resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return nil, err
	}
	getDeposit(deposit_resp, alert)

	return alert, nil

}

func (api *APIClient) GetAssetsTicker(productCode string) (*bitflyer.AssetsTicker, error) {
	url := "ticker"
	resp, err := api.doRequest("GET", url, map[string]string{"product_code": productCode}, nil)
	if err != nil {
		return nil, err
	}
	var assetsticker bitflyer.AssetsTicker
	err = json.Unmarshal(resp, &assetsticker)
	if err != nil {
		return nil, err
	}
	return &assetsticker, nil
}

func (api *APIClient) GetExecution(productCode string) ([]bitflyer.Execution, error) {
	url := "me/getexecutions"
	resp, err := api.doRequest("GET", url, map[string]string{"product_code": productCode}, nil)
	if err != nil {
		return nil, err
	}
	var execution []bitflyer.Execution
	err = json.Unmarshal(resp, &execution)
	if err != nil {
		return nil, err
	}
	return execution, nil
}

func (api *APIClient) CalcProfit(asset float64, jpy float64) (float64, error) {
	url := "me/getwithdrawals"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return 0, err
	}
	var withdrawals []bitflyer.Withdrawals
	err = json.Unmarshal(resp, &withdrawals)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var all_withdrawals float64
	for i := 0; i < len(withdrawals); i++ {
		all_withdrawals += withdrawals[i].Amount
	}

	url = "me/getdeposits"
	deposit_resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return 0, err
	}
	var deposit []bitflyer.Deposits
	err = json.Unmarshal(deposit_resp, &deposit)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var all_deposit float64
	for i := 0; i < len(deposit); i++ {
		all_deposit += deposit[i].Amount
	}

	url = "me/getcoinins"
	coinin_resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return 0, err
	}
	var coinin []bitflyer.Coins
	err = json.Unmarshal(coinin_resp, &coinin)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var all_coinins float64
	for i := 0; i < len(coinin); i++ {
		all_coinins += coinin[i].Amount
	}

	url = "me/getcoinouts"
	coinout_resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return 0, err
	}
	var coinout []bitflyer.Coins
	err = json.Unmarshal(coinout_resp, &coinout)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	url = "ticker"
	var ticker bitflyer.Ticker
	var all_coinouts float64
	for i := 0; i < len(coinout); i++ {
		productCode := coinout[i].CurrencyCode
		productCode = productCode + "_JPY"
		ticker_resp, _ := api.doRequest("GET", url, map[string]string{"product_code": productCode}, nil)
		err = json.Unmarshal(ticker_resp, &ticker)
		if err != nil {
			log.Println(err)
		}
		current_crpto_price := ticker.Ltp
		all_coinouts += current_crpto_price * coinout[i].Amount
	}
	profit := (asset + jpy + all_withdrawals) - (all_deposit)
	return profit, nil
}

func (api *APIClient) SendOrder(order *bitflyer.Order) (*bitflyer.ResponseSendChildOrder, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	url := "me/sendchildorder"
	resp, err := api.doRequest("POST", url, map[string]string{}, data)
	if err != nil {
		return nil, err
	}
	var response bitflyer.ResponseSendChildOrder
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (api *APIClient) ListOrder(query map[string]string) ([]bitflyer.Order, error) {
	resp, err := api.doRequest("GET", "me/getchildorders", query, nil)
	if err != nil {
		return nil, err
	}
	var responseListOrder []bitflyer.Order
	err = json.Unmarshal(resp, &responseListOrder)
	if err != nil {
		return nil, err
	}
	return responseListOrder, nil
}

func (api *APIClient) GetRealTimeTicker(symbol string, ch chan<- Ticker) {
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	channel := fmt.Sprintf("lightning_ticker_%s", symbol)
	type subscribeparams bitflyer.SubscribeParams
	if err := c.WriteJSON(&bitflyer.JsonRPC2{Version: "2.0", Method: "subscribe", Params: &subscribeparams{channel}}); err != nil {
		log.Fatal("subscribe:", err)
		return
	}

OUTER:
	for {
		message := new(bitflyer.JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}

		if message.Method == "channelMessage" {
			switch v := message.Params.(type) {
			case map[string]interface{}:
				for key, binary := range v {
					if key == "message" {
						marshaTic, err := json.Marshal(binary)
						if err != nil {
							continue OUTER
						}
						var ticker Ticker
						if err := json.Unmarshal(marshaTic, &ticker); err != nil {
							continue OUTER
						}
						ch <- ticker
					}
				}
			}
		}
	}
}

type Candle bitflyer.Candle

func NewCandle(productCode string, duration time.Duration, timeDate time.Time, open, close, high, low, volume float64) *Candle {
	return &Candle{
		productCode,
		duration,
		timeDate,
		open,
		close,
		high,
		low,
		volume,
	}
}

func (c *Candle) TableName() string {
	return GetCandleTableName(c.ProductCode, c.Duration)
}

func (c *Candle) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (time, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?)", c.TableName())
	_, err := DbConnection.Exec(cmd, c.Time.Format(time.RFC3339), c.Open, c.Close, c.High, c.Low, c.Volume)
	if err != nil {
		return err
	}
	return err
}

func (c *Candle) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET open = ?, close = ?, high = ?, low = ?, volume = ? WHERE time = ?", c.TableName())
	_, err := DbConnection.Exec(cmd, c.Open, c.Close, c.High, c.Low, c.Volume, c.Time.Format(time.RFC3339))
	if err != nil {
		return err
	}
	return err
}

func GetCandle(productCode string, duration time.Duration, dateTime time.Time) *Candle {
	tableName := GetCandleTableName(productCode, duration)
	cmd := fmt.Sprintf("SELECT time, open, close, high, low, volume FROM %s WHERE time = ?", tableName)
	row := DbConnection.QueryRow(cmd, dateTime.Format(time.RFC3339))
	var candle Candle
	err := row.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume)
	if err != nil {
		return nil
	}
	return NewCandle(productCode, duration, candle.Time, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)
}

func CreateCandleWithDuration(ticker Ticker, productCode string, duration time.Duration) bool {
	currentCandle := GetCandle(productCode, duration, ticker.TruncateDateTime(duration))
	price := ticker.GetMidPrice()
	if currentCandle == nil {
		candle := NewCandle(productCode, duration, ticker.TruncateDateTime(duration),
			price, price, price, price, ticker.Volume)
		candle.Create()
		return true
	}

	if currentCandle.High <= price {
		currentCandle.High = price
	} else if currentCandle.Low >= price {
		currentCandle.Low = price
	}
	currentCandle.Volume += ticker.Volume
	currentCandle.Close = price
	currentCandle.Save()
	return false
}

type DataFrameCandle bitflyer.DataFrameCandle

func (df *DataFrameCandle) Times() []time.Time {
	s := make([]time.Time, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Time
	}
	return s
}

func (df *DataFrameCandle) Opens() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Open
	}
	return s
}

func (df *DataFrameCandle) Closes() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Close
	}
	return s
}

func (df *DataFrameCandle) Highs() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.High
	}
	return s
}

func (df *DataFrameCandle) Lows() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Low
	}
	return s
}

func (df *DataFrameCandle) Volumes() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Volume
	}
	return s
}

func GetAllCandle(productCode string, duration time.Duration, limit int) (dfCandle *DataFrameCandle, err error) {
	tableName := GetCandleTableName(productCode, duration)
	cmd := fmt.Sprintf(`SELECT * FROM (
		SELECT time, open, close, high, low, volume FROM %s ORDER BY time DESC LIMIT ?
		) ORDER BY time ASC;`, tableName)
	rows, err := DbConnection.Query(cmd, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	dfCandle = &DataFrameCandle{}
	dfCandle.ProductCode = productCode
	dfCandle.Duration = duration
	for rows.Next() {
		var candle bitflyer.Candle
		candle.ProductCode = productCode
		candle.Duration = duration
		rows.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume)
		dfCandle.Candles = append(dfCandle.Candles, candle)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return dfCandle, nil
}

type SignalEvent struct {
	ProductCode string    `json:"product_code"`
	Time        time.Time `json:"time"`
	Side        string    `json:"side"`
	Price       float64   `json:"price"`
	Size        float64   `json:"size"`
}

type SignalEvents struct {
	Signals []SignalEvent `json:"signals,omitempty"`
}

func (s *SignalEvent) Save() bool {
	cmd := fmt.Sprintf("INSERT INTO %s(time, product_code, side, price, size) VALUES (?, ?, ?, ?, ?)", tableNameSignalEvents)
	_, err := DbConnection.Exec(cmd, s.Time.Format(time.RFC3339), s.ProductCode, s.Side, s.Price, s.Size)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println(err)
			return true
		}
		return false
	}
	return true
}

func NewSignalEvents() *SignalEvents {
	return &SignalEvents{}
}

func GetSignalEventsByCount(loadEvents int) *bitflyer.SignalEvents {
	cmd := fmt.Sprintf(`SELECT * FROM (
		SELECT time, product_code, side, price, size FROM %s WHERE product_code = ? ORDER BY time DESC LIMIT ? )
		ORDER BY time ASC;`, tableNameSignalEvents)
	rows, err := DbConnection.Query(cmd, config.Config.ProductCode, loadEvents)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var signalEvents bitflyer.SignalEvents
	for rows.Next() {
		var signalEvent bitflyer.SignalEvent
		rows.Scan(&signalEvent.Time, &signalEvent.ProductCode, &signalEvent.Side, &signalEvent.Price, &signalEvent.Size)
		signalEvents.Signals = append(signalEvents.Signals, signalEvent)
	}
	err = rows.Err()
	if err != nil {
		return nil
	}
	return &signalEvents
}

func GetSignalEventsAfterTime(timeTime time.Time) *bitflyer.SignalEvents {
	cmd := fmt.Sprintf(`SELECT * FROM (
		SELECT time, product_code, side, price, size FROM %s 
		WHERE DATETIME(time) >= DATETIME(?)
		ORDER BY time DESC
	) ORDER BY time ASC;`, tableNameSignalEvents)
	rows, err := DbConnection.Query(cmd, timeTime.Format(time.RFC3339))
	if err != nil {
		return nil
	}
	defer rows.Close()

	var signalEvents bitflyer.SignalEvents
	for rows.Next() {
		var signalEvent bitflyer.SignalEvent
		rows.Scan(&signalEvent.Time, &signalEvent.ProductCode, &signalEvent.Side, &signalEvent.Price, &signalEvent.Size)
		signalEvents.Signals = append(signalEvents.Signals, signalEvent)
	}
	err = rows.Err()
	if err != nil {
		return nil
	}
	return &signalEvents
}

func (s *SignalEvents) CanBuy(time time.Time) bool {
	lenSignals := len(s.Signals)
	if lenSignals == 0 {
		return true
	}
	lastSignal := s.Signals[lenSignals-1]
	if lastSignal.Side == "SELL" && lastSignal.Time.Before(time) {
		return true
	}
	return false
}

func (s *SignalEvents) CanSell(time time.Time) bool {
	lenSignals := len(s.Signals)
	if lenSignals == 0 {
		return false
	}
	lastSignal := s.Signals[lenSignals-1]
	if lastSignal.Side == "BUY" && lastSignal.Time.Before(time) {
		return true
	}
	return false
}

func (s *SignalEvents) Buy(ProductCode string, time time.Time, price, size float64, save bool) bool {
	if !s.CanBuy(time) {
		return false
	}
	signalEvent := SignalEvent{
		ProductCode: ProductCode,
		Time:        time,
		Side:        "BUY",
		Price:       price,
		Size:        size,
	}
	if save {
		signalEvent.Save()
	}
	s.Signals = append(s.Signals, signalEvent)
	return true
}

func (s *SignalEvents) Sell(ProductCode string, time time.Time, price, size float64, save bool) bool {
	if !s.CanSell(time) {
		return false
	}
	signalEvent := SignalEvent{
		ProductCode: ProductCode,
		Time:        time,
		Side:        "SELL",
		Price:       price,
		Size:        size,
	}
	if save {
		signalEvent.Save()
	}
	s.Signals = append(s.Signals, signalEvent)
	return true
}

func (s *SignalEvents) Profit() float64 {
	total := 0.0
	beforeSell := 0.0
	isHolding := false

	for i, signalEvent := range s.Signals {
		if i == 0 && signalEvent.Side == "SELL" {
			continue
		}
		if signalEvent.Side == "BUY" {
			total -= signalEvent.Price * signalEvent.Size
			isHolding = true
		}
		if signalEvent.Side == "SELL" {
			total += signalEvent.Price * signalEvent.Size
			isHolding = false
			beforeSell = total
		}
		if isHolding {
			return beforeSell
		}
	}
	return total
}

func (s SignalEvents) MarshalJSON() ([]byte, error) {
	value, err := json.Marshal(&struct {
		Signals []SignalEvent `json:"signals,omitempty"`
		Profit  float64       `json:"profit,omitempty"`
	}{
		Signals: s.Signals,
		Profit:  s.Profit(),
	})
	if err != nil {
		return nil, err
	}
	return value, err
}

func (s *SignalEvents) CollectAfter(time time.Time) *SignalEvents {
	for i, signal := range s.Signals {
		if time.After(signal.Time) {
			continue
		}
		return &SignalEvents{Signals: s.Signals[i:]}
	}
	return nil
}

func (df *DataFrameCandle) AddRsi(period int) bool {
	if len(df.Candles) > period {
		values := talib.Rsi(df.Closes(), period)
		df.Rsi = &bitflyer.Rsi{
			Period: period,
			Values: values,
		}
		return true
	}
	return false
}

func (df *DataFrameCandle) AddEvents(timeTime time.Time) bool {
	signalEvents := GetSignalEventsAfterTime(timeTime)
	if len(signalEvents.Signals) > 0 {
		df.Events = signalEvents
		return true
	}
	return false
}

func AutomaticNotification() {
	for {
		df, _ := GetAllCandle("BTC_JPY", time.Minute, 365)
		df.BackTestRsi(14, 30.0, 70.0)	
	}
}

func (df *DataFrameCandle) BackTestRsi(period int, buyThread, sellThread float64) *SignalEvents {
	lenCandles := len(df.Candles)
	if lenCandles <= period {
		return nil
	}

	SignalEvents := NewSignalEvents()
	values := talib.Rsi(df.Closes(), period)
	for i := 1; i < lenCandles; i++ {
		if values[i-1] == 0 || values[i-1] == 100 {
			continue
		}
		if values[i-1] < buyThread && values[i] > buyThread {
			SignalEvents.Buy(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
			fmt.Println("************************** Buy ***********************")
		}

		if values[i-1] > sellThread && values[i] <= sellThread {
			SignalEvents.Sell(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
			fmt.Println("************************** Sell ***********************")
		}
	}
	return SignalEvents
}

// func (df *DataFrameCandle) OptimizeRsi() (performance float64, bestPeriod int, bestBuyThread, bestSellThread float64) {
// 	bestPeriod = 14
// 	bestBuyThread, bestSellThread = 30.0, 70.0

// 	for period := 5; period < 25; period++ {
// 		SignalEvents := df.BackTestRsi(period, bestSellThread, bestSellThread)
// 		if SignalEvents == nil {
// 			continue
// 		}
// 		profit := SignalEvents.Profit()
// 		if performance < profit {
// 			performance = profit
// 			bestPeriod = period
// 			bestBuyThread = bestBuyThread
// 			bestSellThread = bestSellThread
// 		}
// 	}
// 	return performance, bestPeriod, bestBuyThread, bestSellThread
// }
