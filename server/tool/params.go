package tool

import (
	"os"

	"github.com/joho/godotenv"
)

type Info struct {
	dburl string
}

func (u Info) GetDBUrl() string {
	_ = godotenv.Load()

	postgresURL := os.Getenv("POSTGRES_URL")
	return postgresURL
}
