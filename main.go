package main

import (
	"fmt"

	"github.com/serpedious/automatic-trading-system/bitflyer"
	"github.com/serpedious/automatic-trading-system/config"
	"github.com/serpedious/automatic-trading-system/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}

// package main

// import (
//     "log"
//     "net/http"

//     "github.com/gorilla/mux"
// 	"github.com/gorilla/handlers"
// )

// func public(w http.ResponseWriter, r *http.Request) {
//     w.Write([]byte("hello public!\n"))
// }

// func private(w http.ResponseWriter, r *http.Request) {
//     w.Write([]byte("hello private!\n"))
// }

// func balance(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("this is balance form golang!\n"))
// }

// func main() {
//     allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
//     allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
//     allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

//     r := mux.NewRouter()
//     r.HandleFunc("/public", public)
//     r.HandleFunc("/private", private)
//     r.HandleFunc("/balance", balance)

//     log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
// }
