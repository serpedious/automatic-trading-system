package main

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/config"
	"github.com/serpedious/automatic-trading-system/server/user"
	"github.com/serpedious/automatic-trading-system/server/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func hello() int {
	return 1
}

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

func execution(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	execution_data, _ := apiClient.GetExecution(config.Config.ProductCode)
	js, _ := json.Marshal(execution_data)
	w.Write([]byte(js))
}

func listorder(w http.ResponseWriter, r *http.Request) {
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

func sendorder(w http.ResponseWriter, r *http.Request) {
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

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
			"https://localhost:*",
			"https://api.serpedious.link:*",
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
	r.Get("/execution", execution)
	r.Get("/listorder", listorder)
	r.Get("/verify", user.TokenVerifyMiddleWare(user.VerifyEndpoint))
	r.Post("/signup", user.Signup)
	r.Post("/signin", user.Signin)
	r.Post("/sendorder", sendorder)

	log.Println("server..")
	http.ListenAndServe(":8000", r)
}
