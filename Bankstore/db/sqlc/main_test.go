package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://app_user:pswd@localhost:5432/bankdb?sslmode=disable"
)

// Глобальный контекст для работы с БД и тестами
//var ctx = context.Background()

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	// соединение с бд
	//conn, err := pgx.Connect(ctx, dbSource)
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can not connection to db", err)
	}
	//закрываем соединеие
	defer testDB.Close()
	//Вызваем конструктор создания экзампляра типа Queries
	testQueries = New(testDB)
	//Запускаем subtest(тесты) и итоговый код выполнения передаем в Exit()
	os.Exit(m.Run())
}
