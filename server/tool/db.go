package tool

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	HOST     = os.Getenv("POSTGRES_URL")
	DATABASE = os.Getenv("POSTGRES_DB")
	USER     = os.Getenv("POSTGRES_USER")
	PASSWORD = os.Getenv("PGPASSWORD")
)

var Db *sql.DB

func NewDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		// prod
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", HOST, USER, DATABASE, PASSWORD)
		Db, _ = sql.Open("postgres", dbURI)
		return Db
	}
	// dev
	Db, err := sql.Open("postgres", "host=postgres_db port=5432 user="+USER+" password="+PASSWORD+" dbname="+DATABASE+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return Db
}
