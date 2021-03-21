package tool

import (
	// "database/sql"
	// "log"
	"os"

	// _ "github.com/lib/pq"
)

var (
	HOST     = os.Getenv("POSTGRES_URL")
	DATABASE = os.Getenv("POSTGRES_DB")
	USER     = os.Getenv("POSTGRES_USER")
	PASSWORD = os.Getenv("PGPASSWORD")
)
// var Db *sql.DB

func Connect() {
	// var (
	// 	HOST     = os.Getenv("POSTGRES_URL")
	// 	DATABASE = os.Getenv("POSTGRES_DB")
	// 	USER     = os.Getenv("POSTGRES_USER")
	// 	PASSWORD = os.Getenv("PGPASSWORD")
	// )
	// i := Info{}
	// pgUrl, err := pq.ParseURL(i.GetDBUrl())
	// if err != nil {
	// 	log.Fatal()
	// }

	// neet to implement db connection in the dev env
	// Db, err := sql.Open("postgres", "host="+HOST+" port=5432 user="+USER+" password="+PASSWORD+" dbname="+DATABASE+" sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = Db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
