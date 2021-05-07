package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	ApiKey             string
	ApiSecret          string
	LogFile            string
	ProductCode        string
	CoinCheckApiKey    string
	CoinCheckApiSecret string
	DMMApiKey          string
	DMMApiSecret       string
	SBIApiKey          string
	SBIApiSecret       string
	TradeDuration      time.Duration
	Durations          map[string]time.Duration
	DbName             string
	SQLDriver          string
	Port               int
}

var Config ConfigList

func init() {
	_ = godotenv.Load()

	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	Config = ConfigList{
		ApiKey:      os.Getenv("BITFLYER_API_KEY"),
		ApiSecret:   os.Getenv("BITFLYER_API_KEY_SECRET"),
		LogFile:     os.Getenv("LOG_FILE"),
		ProductCode: os.Getenv("PRODUCT_CODE"),

		CoinCheckApiKey:    os.Getenv("COINCHECK_API_KEY"),
		CoinCheckApiSecret: os.Getenv("COINCHECK_API_KEY_SECRET"),
		DMMApiKey:          os.Getenv("DMM_API_KEY"),
		DMMApiSecret:       os.Getenv("DMM_API_KEY_SECRET"),
		SBIApiKey:          os.Getenv("SBI_API_KEY"),
		SBIApiSecret:       os.Getenv("SBI_API_KEY_SECRET"),
		Durations:          durations,
		TradeDuration:      durations[os.Getenv("TRADE_DURATION")],
		DbName:             os.Getenv("DB_NAME"),
		SQLDriver:          os.Getenv("SQL_DRIVER"),
	}
}

// package config

// import (
// 	"log"
// 	"os"

// 	"gopkg.in/ini.v1"
// )

// type ConfigList struct {
// 	ApiKey      string
// 	ApiSecret   string
// 	LogFile     string
// 	ProductCode string
// }

// var Config ConfigList

// func init() {
// 	cfg, err := ini.Load("config.ini")
// 	if err != nil {
// 		log.Printf("Failed to read config file: %v", err)
// 		os.Exit(1)
// 	}

// 	Config = ConfigList{
// 		ApiKey:      cfg.Section("bitflyer").Key("api_key").String(),
// 		ApiSecret:   cfg.Section("bitflyer").Key("api_secret").String(),
// 		LogFile:     cfg.Section("gotrading").Key("log_file").String(),
// 		ProductCode: cfg.Section("gotrading").Key("product_code").String(),
// 	}
// }
