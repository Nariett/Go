package main

import (
	"Server/internal/server"

	_ "github.com/lib/pq"
)

func main() {
	server.StartServer()
}
