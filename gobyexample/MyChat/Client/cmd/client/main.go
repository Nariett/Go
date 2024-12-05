package main

import (
	pb "MyChat/proto"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	dbLocalhost := "localhost:" + os.Getenv("DB_LOCALHOST")
	conn, err := grpc.Dial(dbLocalhost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	var (
		name     string
		password string
		flag     bool = false
	)
	var value int
	for {
		fmt.Println("1 - Войти в чат\n2 - Зарегистрироваться в чате\n3 - Выйти из чата")
		fmt.Scanln(&value)
		switch value {
		case 1:
			fmt.Println("Введите имя: ")
			fmt.Scanln(&name)
			fmt.Println("Введите пароль: ")
			fmt.Scanln(&password)
			response, err := client.AuthUser(context.Background(), &pb.UserData{Name: name, Password: password})
			if err != nil {
				log.Fatalf("Ошибка аутентификации: %v", err)
			}
			if response.Success {
				fmt.Println("Вы вошли в систему!")
				flag = true
			} else {
				fmt.Println(response.Message)
			}
		case 2:
			for {
				fmt.Println("Введите имя: ")
				fmt.Scanln(&name)
				fmt.Println("Введите пароль: ")
				fmt.Scanln(&password)
				response, err := client.RegUser(context.Background(), &pb.UserData{Name: name, Password: password})
				if err != nil {
					log.Fatalf("Ошибка регистрации: %v", err)
				}
				if response.Success {
					fmt.Println("Вы прошли регистрацию!")
					break
				} else {
					fmt.Println(response.Message)
				}
			}

		case 3:
			fmt.Println("Вы вышли из чата...")
			os.Exit(1)
		}
		if flag {
			break
		}
	}

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
			fmt.Printf("Новое сообщение от %s: %s\n", msg.Sender, msg.Content)
		}
	}()
	for {
		var recipient, msg string

		fmt.Println("Введите имя, кому хотите отправить сообщение: ")
		fmt.Scanln(&recipient)

		fmt.Println("Введите сообщение: ")
		fmt.Scanln(&msg)

		if len(recipient) != 0 && len(msg) != 0 {
			_, err := client.SendMessage(context.Background(), &pb.UserMessage{
				Sender:    name,
				Recipient: recipient,
				Content:   msg,
			})
			if err != nil {
				log.Printf("Ошибка отправки сообщения: %v", err)
			}
		} else {
			fmt.Println("Сообщение не отправлено. Введите имя пользователя и сообщение.")
		}
	}
}
