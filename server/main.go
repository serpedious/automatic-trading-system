package main

import (
	"log"

	"net/http"

	"github.com/serpedious/automatic-trading-system/server/config"
	"github.com/serpedious/automatic-trading-system/server/controller"
	"github.com/serpedious/automatic-trading-system/server/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func hello() int {
	return 1
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

	r.Get("/verify", controller.TokenVerifyMiddleWare(controller.VerifyEndpoint))
	r.Post("/signup", controller.Signup)
	r.Post("/signin", controller.Signin)

	r.Post("/sendorder", controller.SendOrder)
	r.Get("/balance", controller.Balance)
	r.Get("/ticker", controller.Ticker)
	r.Get("/execution", controller.GetExecution)
	r.Get("/listorder", controller.Listorder)

	r.Get("/getallmemos", controller.GetAllMemos)
	r.Post("/creatememo", controller.CreateMemo)
	r.Put("/donememo", controller.DoneMemo)
	r.Put("/deletememo", controller.DeleteMemo)

	log.Println("server...")
	http.ListenAndServe(":8000", r)
}
