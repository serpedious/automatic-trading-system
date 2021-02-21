package tool

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Info struct {
	dburl string
}

func (u Info) GetDBUrl() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("production...")	
	}
	postgresURL := os.Getenv("POSTGRES_URL")
	return postgresURL
}
