package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
}

func BuildConnStr() string {
	LoadEnv()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)
}

func GetProtocolAndPort() (protocol, port string) {
	LoadEnv()
	dbProtocol := os.Getenv("PROTOCOL")
	dbPort := os.Getenv("DB_PORT")
	return dbProtocol, dbPort
}
