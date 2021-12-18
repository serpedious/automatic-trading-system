package bitflyer

import "time"

const BaseURL = "https://api.bitflyer.com/v1/"

type Balance struct {
	CurrentCode string  `json:"currency_code"`
	Available   float64 `json:"available"`
}

type BalanceHistory struct {
	ID           int     `json:"id"`
	TradeDate    string  `json:"trade_date"`
	EventDate    string  `json:"event_date"`
	ProductCode  string  `json:"product_code"`
	CurrencyCode string  `json:"currency_code"`
	TradeType    string  `json:"trade_type"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Quantity     float64 `json:"quantity"`
	Commission   float64 `json:"commission"`
	Balance      float64 `json:"balance"`
	OrderID      string  `json:"order_id"`
}

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

type AssetsTicker struct {
	ProductCode string  `json:"product_code"`
	Ltp         float64 `json:"ltp"`
}

type Execution struct {
	ID                     int     `json:"id"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	ExecDate               string  `json:"exec_date"`
	ChildOrderID           string  `json:"child_order_id"`
	Commission             float64 `json:"commission"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
}

type Order struct {
	ID                     int     `json:"id"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	ProductCode            string  `json:"product_code"`
	ChildOrderType         string  `json:"child_order_type"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	MinuteToExpires        int     `json:"minute_to_expire"`
	TimeInForce            string  `json:"time_in_force"`
	Status                 string  `json:"status"`
	ErrorMessage           string  `json:"error_message"`
	AveragePrice           float64 `json:"average_price"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	OutstandingSize        float64 `json:"outstanding_size"`
	CancelSize             float64 `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommission        float64 `json:"total_commission"`
	Count                  int     `json:"count"`
	Before                 int     `json:"before"`
	After                  int     `json:"after"`
}

type ResponseSendChildOrder struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

type JsonRPC2 struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result,omitempty"`
	Id      *int        `json:"id,omitempty"`
}

type SubscribeParams struct {
	Channel string `json:"channel"`
}

type Candle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Time        time.Time     `json:"time"`
	Open        float64       `json:"open"`
	Close       float64       `json:"close"`
	High        float64       `json:"high"`
	Low         float64       `json:"low"`
	Volume      float64       `json:"volume"`
}

type DataFrameCandle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Candles     []Candle      `json:"candles"`
}

type MyAssets struct {
	Crpto  string  `json:"crpto"`
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	Value  int     `json:"value"`
}

type Trade struct {
	Price float64 `json:"price"`
}

type Deposits struct {
	ID           int    `json:"id"`
	OrderID      string `json:"order_id"`
	CurrencyCode string `json:"currency_code"`
	Amount       float64    `json:"amount"`
	Status       string `json:"status"`
	EventDate    string `json:"event_date"`
}

type Withdrawals struct {
	ID           int    `json:"id"`
	OrderID      string `json:"order_id"`
	CurrencyCode string `json:"currency_code"`
	Amount       float64    `json:"amount"`
	Status       string `json:"status"`
	EventDate    string `json:"event_date"`
}

type Coins struct {
	ID            int     `json:"id"`
	OrderID       string  `json:"order_id"`
	CurrencyCode  string  `json:"currency_code"`
	Amount        float64 `json:"amount"`
	Address       string  `json:"address"`
	TxHash        string  `json:"tx_hash"`
	Fee           float64 `json:"fee"`
	AdditionalFee float64 `json:"additional_fee"`
	Status        string  `json:"status"`
	EventDate     string  `json:"event_date"`
}
