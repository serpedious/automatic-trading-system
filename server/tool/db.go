package tool

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connect() {
	i := Info{}
	pgUrl, err := pq.ParseURL(i.GetDBUrl())
	if err != nil {
		log.Fatal()
	}

	// Db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=docker dbname=test_db sslmode=disable")

	Db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
