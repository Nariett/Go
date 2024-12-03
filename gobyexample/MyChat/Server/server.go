package main

import (
	pb "chat/chat"
	"context"
	"log"
	"net"
	"strings"
	"sync"

	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	mu    sync.Mutex
	users map[string]chan pb.UserMessage
}

func newChatServer() *ChatServer {
	return &ChatServer{
		users: make(map[string]chan pb.UserMessage),
	}
}

func (c *ChatServer) JoinChat(user *pb.User, stream pb.ChatService_JoinChatServer) error {
	c.mu.Lock()
	msgChan := make(chan pb.UserMessage, 10)
	c.users[user.Name] = msgChan
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		delete(c.users, user.Name)
		close(msgChan)
		c.mu.Unlock()
	}()

	for msg := range msgChan {
		if err := stream.Send(&msg); err != nil {
			return err
		}
	}
	return nil

}
func (c *ChatServer) GetUsers(ctx context.Context, user *pb.User) (*pb.Users, error) {

	c.mu.Lock()
	var activeUsers []string
	for key := range c.users {
		activeUsers = append(activeUsers, key)
	}
	c.mu.Unlock()
	result := strings.Join(activeUsers, ", ")
	usersRespons := &pb.Users{
		Usernames: result,
	}
	return usersRespons, nil
}
func (c *ChatServer) SendMessage(ctx context.Context, msg *pb.UserMessage) (*pb.Empty, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.users[msg.Recipient]; exists {
		ch <- *msg
	}
	return &pb.Empty{}, nil

}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка сервера: %v\n", err)
	}
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, newChatServer())
	log.Println("Сервер запущен на порту: 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Ошибка сервера: %v\n", err)
	}

}
