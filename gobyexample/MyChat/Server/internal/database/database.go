package database

import (
	pb "MyChat/proto"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func buildConnStr() string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)
}
func RegUser(user *pb.UserData) (*pb.ServerResponse, error) {
	db, err := sql.Open("postgres", buildConnStr())
	if err != nil {
		log.Fatalf("Ошибка поключение к базе данных: %v\n", err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Users (name, password) values ($1, $2)", user.Name, user.Password)
	if err != nil {
		log.Printf("Данный ник уже занят: %v\n", err)
		return &pb.ServerResponse{
			Success: false,
			Message: "Пользователь не добавлен в базу данных, так как ник уже занят",
		}, nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Ошибка получения последнего ID: %v", err)
	}
	log.Printf("Добавлен новый пользователь: id: %d, name: %s, password: %s\n", id, user.Name, user.Password)
	return &pb.ServerResponse{
		Success: true,
		Message: "Пользователь добавлен в базу данных",
	}, nil
}
func AuthUser(user *pb.UserData) (*pb.ServerResponse, error) {
	db, err := sql.Open("postgres", buildConnStr())
	if err != nil {
		log.Fatalf("Ошибка базы данных: %v\n", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Users where name = $1 and password = $2", user.Name, user.Password)
	if err != nil {
		log.Fatalf("Ошибка получения данных: %v\n", err)
	}
	if !rows.Next() {
		log.Printf("Данный пользователь не найден: name:%s password:%s ", user.Name, user.Password)
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
