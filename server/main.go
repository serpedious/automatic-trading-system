package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/serpedious/automatic-trading-system/server/bitflyer"
	"github.com/serpedious/automatic-trading-system/server/config"
	"github.com/serpedious/automatic-trading-system/server/tool"
	"github.com/serpedious/automatic-trading-system/server/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

func verifyEndpoint(w http.ResponseWriter, r *http.Request) {
	responseByJSON(w, "authentication is ok!")
}

func tokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var errorObject Error

		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		fmt.Println("bearerToken: ", bearerToken)

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("An error has occured")
				}
				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				errorInResponse(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				errorInResponse(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "This token is invalid"
			return
		}
	})
}

func createToken(user User) (string, error) {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "__init__",
	})

	spew.Dump(token)
	tokenString, err := token.SignedString([]byte(secret))

	fmt.Println("--------------------")
	fmt.Println("tokenString: ", tokenString)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func errorInResponse(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
	return
}

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello private!\n"))
}

func responseByJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
	return
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error

	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is required"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is required"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	fmt.Println(user)
	fmt.Println("------------------------")
	spew.Dump(user)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("password: ", user.Password)
	fmt.Println("hash password: ", hash)

	user.Password = string(hash)
	fmt.Println("converted password: ", user.Password)

	sql_query := "INSERT INTO USERS(EMAIL, PASSWORD) VALUES($1, $2) RETURNING ID;"

	err = db.QueryRow(sql_query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		error.Message = "server error"
		errorInResponse(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")

	responseByJSON(w, user)
}

func login(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error
	var jwt JWT

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is required"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is required"
		errorInResponse(w, http.StatusBadRequest, error)
	}
	password := user.Password
	fmt.Println("password: ", password)

	row := db.QueryRow("SELECT * FROM USERS WHERE email=$1;", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "no existed user"
			errorInResponse(w, http.StatusBadRequest, error)
		} else {
			log.Fatal(err)
		}
	}

	hasedPassword := user.Password
	fmt.Println("hasedPassword: ", hasedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))

	if err != nil {
		error.Message = "This is invalid password"
		errorInResponse(w, http.StatusUnauthorized, error)
		return
	}

	token, err := createToken(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	responseByJSON(w, jwt)
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
	i := tool.Info{}
	pgUrl, err := pq.ParseURL(i.GetDBUrl())
	if err != nil {
		log.Fatal()
	}

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

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
	r.Get("/verify", tokenVerifyMiddleWare(verifyEndpoint))
	r.Post("/signup", signup)
	r.Post("/login", login)

	log.Println("server..")
	http.ListenAndServe(":8000", r)
}
