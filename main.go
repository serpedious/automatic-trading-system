package main

import (
	"fmt"

	"net/http"

	"github.com/serpedious/automatic-trading-system/bitflyer"
	"github.com/serpedious/automatic-trading-system/config"
	"github.com/serpedious/automatic-trading-system/utils"

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
	w.Write([]byte("this is balance form golang!\n"))
}

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
	
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

	http.ListenAndServe(":8000", r)
}
