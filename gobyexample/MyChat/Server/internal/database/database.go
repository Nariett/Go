package database

import (
	"database/sql"
	"log"

	proto "github.com/Nariett/Go/gobyexample/MyChat/Proto"
)

func RegUser(db *sql.DB, user *proto.UserData) (*proto.ServerResponse, error) {
	result, err := db.Exec("insert into Users (name, password) values ($1, $2)", user.Name, user.Password)
	if err != nil {
		log.Printf("Данный ник уже занят: %v\n", err)
		return &proto.ServerResponse{
			Success: false,
			Message: "Пользователь не добавлен в базу данных, так как ник уже занят",
		}, nil
	}
	id, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	log.Printf("Добавлен новый пользователь: id: %d, name: %s, password: %s\n", id, user.Name, user.Password)
	return &proto.ServerResponse{
		Success: true,
		Message: "Пользователь добавлен в базу данных",
	}, nil
}
func AuthUser(db *sql.DB, user *proto.UserData) (*proto.ServerResponse, error) {
	log.Printf("Найти пользователя %s, %s", user.Name, user.Password) //////
	rows, err := db.Query("select * from Users where name = $1 and password = $2", user.Name, user.Password)
	if err != nil {
		log.Fatalf("Ошибка получения данных: %v\n", err)
	}
	defer rows.Close()

	if !rows.Next() {
		log.Printf("Данный пользователь не найден: name:%s password:%s ", user.Name, user.Password)
		return &proto.ServerResponse{
			Success: false,
			Message: "Данный пользователь не найден, повторите попытку.",
		}, nil
	}
	return &proto.ServerResponse{Success: true}, nil
}
