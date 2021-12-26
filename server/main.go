package main

import (
	"log"

	"net/http"

	"github.com/serpedious/automatic-trading-system/server/bitflyer/usecase"
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

	// df, _ := usecase.GetAllCandle("BTC_JPY", time.Minute, 365)
	go usecase.AutomaticNotification()
	// c1 := df.Candles[len(df.Candles)-2]
	// c2 := df.Candles[len(df.Candles)-1]
	// s.Buy("BTC_JPY", c1.Time.UTC(), c1.Close, 1.0, true)
	// s.Sell("BTC_JPY", c2.Time.UTC(), c2.Close, 1.0, true)

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
	r.Get("/getalert", controllers.GetAlert)
	r.Get("/getmyassets", controllers.GetMyAssets)
	r.Get("/getbalancehistory", controllers.GetBalanceHistory)
	r.Get("/execution", controllers.GetExecution)
	r.Get("/listorder", controllers.Listorder)
	r.Get("/calcprofit", controllers.CalcProfit)

	r.Get("/getallmemos", controllers.GetAllMemos)
	r.Post("/creatememo", controllers.CreateMemo)
	r.Put("/donememo", controllers.DoneMemo)
	r.Put("/deletememo", controllers.DeleteMemo)

	r.Get("/getcsvfile", controllers.GetCsvFile)

	r.Get("/compareticker", controllers.CompareTicker)
	r.Get("/api/candle/", controllers.ApiMakeHandler(controllers.ApiCandleHandler))

	log.Println("server...")
	http.ListenAndServe(":8000", r)
}
