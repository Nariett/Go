package server

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func buildConnStr() string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)
}

func getProtocolAndPort() (protocol, port string) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	dbProtocol := os.Getenv("PROTOCOL")
	dbPort := os.Getenv("DB_PORT")
	return dbProtocol, dbPort
}
