package main

import (
	pb "chat/chat"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"sync"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const connStr = "user=postgres password=1111 dbname=ChatDB sslmode=disable"

type chatUser struct {
	id       int32
	name     string
	password string
}

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
func (c *ChatServer) GetUsers(ctx context.Context, user *pb.User) (*pb.ActiveUsers, error) {

	c.mu.Lock()
	var activeUsers []string
	for key := range c.users {
		activeUsers = append(activeUsers, key)
	}
	c.mu.Unlock()
	usersRespons := &pb.ActiveUsers{
		Usernames: activeUsers,
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

func (c *ChatServer) RegUser(ctx context.Context, user *pb.UserData) (*pb.ServerResponse, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Ошибка базы данных: %v\n", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Users where name = $1 and password = $2", user.Name, user.Password)
	if err != nil {
		log.Printf("Ошибка получения данных: %v\n", err)
	}
	if !rows.Next() {
		log.Printf("Нет данных для заданного имени и пароля.")

		result, err := db.Exec("insert into Users (name, password) values ($1, $2)", user.Name, user.Password)
		if err != nil {
			panic(err)
		}
		log.Println(result.RowsAffected())
		return &pb.ServerResponse{
			Success: true,
			Message: "Пользователь добавлен в базу данных",
		}, nil
	}
	return &pb.ServerResponse{
		Success: false,
		Message: "Пользователь не добавлен в базу данных. Придумайте другой логин или пароль",
	}, nil
}

func (c *ChatServer) AuthUser(ctx context.Context, user *pb.UserData) (*pb.ServerResponse, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Ошибка базы данных: %v\n", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Users where name = $1 and password = $2", user.Name, user.Password)
	if err != nil {
		log.Printf("Ошибка получения данных: %v\n", err)
	}
	if !rows.Next() {
		log.Printf("Данный пользователь не найден. Повторите попытку.")
		return &pb.ServerResponse{
			Success: false,
			Message: "Данный пользователь не найден, повторите попытку.",
		}, nil
	} else {
		return &pb.ServerResponse{
			Success: true,
		}, nil
	}
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
