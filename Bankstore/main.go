package main

import (
	"Bankstore/api"
	db "Bankstore/db/sqlc"
	"context"
	"log"

	//"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://app_user:pswd@localhost:5432/bankdb?sslmode=disable"
	serverAddrass = "0.0.0.0:8080"
)

func main() {
	// соединение с бд
	var err error
	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can not connection to db", err)
	}
	//закрываем соединеие после обработки
	defer pool.Close()
	// Создаем еобходимые эземплярыдля работы
	store := db.NewStore(pool)     // Хранилище
	server := api.NewServer(store) // роутинг и прочее

	err = server.Start(serverAddrass)
	if err != nil {
		log.Fatal("can not start server", err)
	}
}
