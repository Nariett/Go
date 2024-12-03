package main

import (
	"bufio"
	pb "chat/chat"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите имя: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	stream, err := client.JoinChat(context.Background(), &pb.User{Name: name})
	if err != nil {
		log.Fatalf("Ошибка подключения к чату: %v", err)
	}

	users, err := client.GetUsers(context.Background(), &pb.User{Name: name})
	if err != nil {
		log.Fatalf("Ошибка получения списка пользователй: %v", err)
	}
	fmt.Println("Список всех пользователей:", users.Usernames)
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Ошибка получения сообщения: %v", err)
			}
			fmt.Printf("Новое сообщение от %s: %s\n", msg.Sender, msg.Message)
		}
	}()
	for {
		var recipient, msg string

		fmt.Println("Введите имя, кому хотите отправить сообщение: ")
		fmt.Scanln(&recipient)

		fmt.Println("Введите сообщение: ")
		fmt.Scanln(&msg)

		_, err := client.SendMessage(context.Background(), &pb.UserMessage{
			Sender:    name,
			Recipient: recipient,
			Message:   msg,
		})
		if err != nil {
			log.Printf("Ошибка отправки сообщения: %v", err)
		}
	}
}
