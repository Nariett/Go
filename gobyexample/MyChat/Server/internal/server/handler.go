package server

import (
	pb "MyChat/proto"
	"Server/internal/database"
	"context"
	"log"
	"strings"
	"sync"
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
			log.Printf("Ошибка отправки сообщения клиенту %s: %v", user.Name, err)
			return err
		}
	}
	return nil
}

func (c *ChatServer) GetUsers(ctx context.Context, user *pb.User) (*pb.ActiveUsers, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var activeUsers []string
	for key := range c.users {
		activeUsers = append(activeUsers, key)
	}
	log.Println("Активные пользователи:", strings.Join(activeUsers, " "))
	return &pb.ActiveUsers{Usernames: activeUsers}, nil
}

func (c *ChatServer) SendMessage(ctx context.Context, msg *pb.UserMessage) (*pb.Empty, error) {
	go func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		if ch, exists := c.users[msg.Recipient]; exists {
			ch <- *msg
		}
	}()
	return &pb.Empty{}, nil
}

func (c *ChatServer) RegUser(ctx context.Context, user *pb.UserData) (*pb.ServerResponse, error) {
	resultChan := make(chan *pb.ServerResponse)
	errorChan := make(chan error)
	go func() {
		responce, err := database.RegUser(user)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- responce
	}()
	select {
	case responce := <-resultChan:
		return responce, nil
	case err := <-errorChan:
		log.Printf("Ошибка при регистрации: %v", err)
		return nil, err
	case <-ctx.Done():
		log.Printf("Контекст завершен: %v", ctx.Err())
		return nil, ctx.Err()
	}
}

func (c *ChatServer) AuthUser(ctx context.Context, user *pb.UserData) (*pb.ServerResponse, error) {
	resultChan := make(chan *pb.ServerResponse)
	errorChan := make(chan error)
	go func() {
		responce, err := database.AuthUser(user)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- responce
	}()
	select {
	case responce := <-resultChan:
		return responce, nil
	case err := <-errorChan:
		log.Printf("Ошибка авторизации: %v", err)
		return nil, err
	case <-ctx.Done():
		log.Printf("Контекст завершен: %v", ctx.Err())
		return nil, ctx.Err()
	}
}
