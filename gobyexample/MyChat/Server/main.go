package main

import (
	"Server/config"
	"Server/internal/server"
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
)

func main() {
	config := config.LoadConfig()

	connStr := config.BuildConnStr()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	protocol, port := config.GetProtocolAndPort()
	listener, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	log.Printf("Сервер запущен на порту: %s", port)
	server.StartServer(listener, db)
}
