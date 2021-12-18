package main

import (
	"log"

	"net/http"

	"github.com/serpedious/automatic-trading-system/server/config"

	"github.com/serpedious/automatic-trading-system/server/controllers"
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

	go controllers.StreamIngectionData()

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
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Set-Cookie"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
	}).Handler)

	r.Get("/verify", controllers.TokenVerifyMiddleWare(controllers.VerifyEndpoint))
	r.Post("/signup", controllers.Signup)
	r.Post("/signin", controllers.Signin)
	r.Get("/signout", controllers.SignOut)

	r.Get("/getuser", controllers.GetUser)
	r.Post("/editpass", controllers.EditPass)

	r.Post("/sendorder", controllers.SendOrder)
	r.Get("/balance", controllers.Balance)
	r.Get("/ticker", controllers.Ticker)
	r.Get("/getmyassets", controllers.GetMyAssets)
	r.Get("/getbalancehistory", controllers.GetBalanceHistory)
	r.Get("/execution", controllers.GetExecution)
	r.Get("/listorder", controllers.Listorder)
	r.Get("/calcprofit", controllers.CalcProfit)

	r.Get("/getallmemos", controllers.GetAllMemos)
	r.Post("/creatememo", controllers.CreateMemo)
	r.Put("/donememo", controllers.DoneMemo)
	r.Put("/deletememo", controllers.DeleteMemo)

	r.Get("/compareticker", controllers.CompareTicker)
	r.Get("/api/candle/", controllers.ApiMakeHandler(controllers.ApiCandleHandler))

	log.Println("server...")
	http.ListenAndServe(":8000", r)
}
