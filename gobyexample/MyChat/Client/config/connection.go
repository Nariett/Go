package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConnStr() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	return os.Getenv("DB_HOST") + os.Getenv("DB_PORT")
}
