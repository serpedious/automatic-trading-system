package tool

import (
	"os"

	"github.com/joho/godotenv"
)

type Info struct {
	dburl string
}

func (u Info) GetDBUrl() string {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	postgresURL := os.Getenv("POSTGRES_URL")
	return postgresURL
}
