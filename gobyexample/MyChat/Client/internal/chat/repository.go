package chat

import (
	pb "MyChat/proto"
	"context"
)

type ChatRepository struct {
	client pb.ChatServiceClient
}

func NewChatRepository(client pb.ChatServiceClient) *ChatRepository {
	return &ChatRepository{client: client}
}

func (r *ChatRepository) AuthUser(name, password string) (*pb.ServerResponse, error) {
	return r.client.AuthUser(context.Background(), &pb.UserData{Name: name, Password: password})
}

func (r *ChatRepository) RegUser(name, password string) (*pb.ServerResponse, error) {
	return r.client.RegUser(context.Background(), &pb.UserData{Name: name, Password: password})
}
