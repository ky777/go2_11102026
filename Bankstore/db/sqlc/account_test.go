/*
Pfttern: <filename>_test.go
*/
package db

import (
	"Bankstore/utils"
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

const (
	dbSource = "postgresql://app_user:pswd@localhost:5432/bankdb?sslmode=disable"
)

// Глобальный контекст для работы с БД и тестами
var ctx = context.Background()
var testQueries *Queries

func TestMain(m *testing.M) {
	// соединение с бд
	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("can not connection to db", err)
	}
	//закрываем соединеие
	defer conn.Close(ctx)
	//Вызваем конструктор создания экзампляра типа Queries
	testQueries = New(conn)
	//Запускаем subtest(тесты) и итоговый код выполнения передаем в Exit()
	os.Exit(m.Run())
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) Account {
	ra := utils.RandomAccount()
	arg := CreateAccountParams{
		Owner:    ra.Owner,
		Balance:  ra.Balance,
		Currency: Currency(ra.Currency),
	}
	account, err := testQueries.CreateAccount(ctx, arg)
	// Две проверки на результат работы CreateAccount
	require.NoError(t, err)
	require.NotEmpty(t, account)
	//Дополнительные проверки
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
