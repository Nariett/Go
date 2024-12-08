package server

import (
	pb "MyChat/proto"
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer(listener net.Listener, db *sql.DB) {
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, newChatServer(db))

	log.Println("gRPC-сервер запущен")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Ошибка запуска gRPC-сервера: %v", err)
	}
}
