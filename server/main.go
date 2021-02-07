package main

import (
	"encoding/json"

	"net/http"

	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/config"
	"github.com/serpedious/automatic-trading-system/server/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello private!\n"))
}

func balance(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	balance_data, _ := apiClient.GetBalance()
	js, _ := json.Marshal(balance_data)
	w.Write([]byte(js))
}

func ticker(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker_data, _ := apiClient.GetTicker(config.Config.ProductCode)
	js, _ := json.Marshal(ticker_data)
	w.Write([]byte(js))
}

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
			"https://localhost:*",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
	}).Handler)

	r.Get("/public", public)
	r.Get("/private", private)
	r.Get("/balance", balance)
	r.Get("/ticker", ticker)

	http.ListenAndServe(":8000", r)
}
