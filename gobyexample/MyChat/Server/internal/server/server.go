package server

import (
	pb "MyChat/proto"
	"Server/config"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer() {
	protocol, dbPort := config.GetProtocolAndPort()
	listener, err := net.Listen(protocol, dbPort)
	if err != nil {
		log.Fatalf("Ошибка сервера: %v\n", err)
	}
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, newChatServer())
	log.Printf("Сервер запущен на порту: %s", dbPort)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Ошибка сервера: %v\n", err)
	}
}
