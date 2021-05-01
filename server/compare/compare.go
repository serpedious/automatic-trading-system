package compare

const (
	BaseURLForCoinCheck = "https://api.coincheck.com/v1/"
	BaseURLForDMM       = "https://api.dmm.com/v1"
	BaseURLForSBI       = "https://api.sbi.com/v1"
)

type CoinCheckTicker struct {
	Last      float64 `json:"last"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Timestamp int     `json:"timestamp"`
}
