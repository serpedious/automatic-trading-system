package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	ApiKey      string
	ApiSecret   string
	LogFile     string
	ProductCode string
}

var Config ConfigList

func init() {
	_ = godotenv.Load()

	Config = ConfigList{
		ApiKey:      os.Getenv("BITFLYER_API_KEY"),
		ApiSecret:   os.Getenv("BITFLYER_API_KEY_SECRET"),
		LogFile:     os.Getenv("LOG_FILE"),
		ProductCode: os.Getenv("PRODUCT_CODE"),
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
