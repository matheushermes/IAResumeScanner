package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_PORT 	= ""
)

func LoadingEnvironmentVariables() {
	var err error

	if err = godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	API_PORT = os.Getenv("API_PORT")
	if err != nil {
		log.Fatal("Error in loading environment variable ", err)
	}
}
