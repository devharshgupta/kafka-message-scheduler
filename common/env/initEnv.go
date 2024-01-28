package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {

	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "local")
	}

	if err := godotenv.Load(fmt.Sprintf("./.env.%s", os.Getenv("APP_ENV"))); err != nil {
		log.Fatal("Error loading .env file")
	}
}
