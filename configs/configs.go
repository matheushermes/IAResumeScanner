package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_PORT 				=	""
	UNIDOC_LICENSE_API_KEY 	=	""
)

func LoadingEnvironmentVariables() {
	var err error

	if err = godotenv.Load("../.env"); err != nil {
		log.Fatal("Erro ao carregar .env ", err)
	}

	API_PORT = os.Getenv("API_PORT")
	UNIDOC_LICENSE_API_KEY = os.Getenv("UNIDOC_LICENSE_API_KEY")
	
	if err != nil {
		log.Fatal("Error em carregar variável de ambiente ", err)
	}
}
