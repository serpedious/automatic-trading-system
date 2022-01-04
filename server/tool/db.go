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
	DATABASE = os.Getenv("POSTGRES_DB")
	USER     = os.Getenv("POSTGRES_USER")
	PASSWORD = os.Getenv("PGPASSWORD")

	HOST          = os.Getenv("POSTGRES_URL")
	PROD_DATABASE = os.Getenv("PROD_POSTGRES_DB")
	PROD_USER     = os.Getenv("PROD_POSTGRES_USER")
	PROD_PASSWORD = os.Getenv("PROD_PGPASSWORD")
)

var Db *sql.DB

func NewDb() *sql.DB {
	PROD := os.Getenv("DB_PROD")
	// distinguish whether env-file exsists or not
	err := godotenv.Load()
	fmt.Println(err)

	if err != nil || PROD == "true" {
		// prod
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", HOST, PROD_USER, PROD_DATABASE, PROD_PASSWORD)
		Db, err = sql.Open("postgres", dbURI)
		if err != nil {
			log.Fatalln(err)
		}
		return Db
	}
	// dev
	Db, err := sql.Open("postgres", "host=postgres_db port=5432 user="+USER+" password="+PASSWORD+" dbname="+DATABASE+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return Db
}
