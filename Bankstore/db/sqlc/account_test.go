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
func TestGetAccount(t *testing.T) {
	account := createRandomAccountSimple(t)
	a, err := testQueries.GetAccount(ctx, account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, a)

	require.Equal(t, account.ID, a.ID, "Идентификаторы должны быть равны.")
	require.Equal(t, account.Owner, a.Owner, "Владелцы должны быть равны.")
	require.Equal(t, account.Balance, a.Balance, "Балансы должны быть равны.")
	require.Equal(t, account.Currency, a.Currency, "Валюты должны быть равны.")
	require.Equal(t, account.CreatedAt, a.CreatedAt, "Даты создания должны быть равны.")
}

func TestListAccounts(t *testing.T) {

	createRandomAccountSimple(t)
	arg := ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}
	ls, err := testQueries.ListAccounts(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, ls)
	require.NotEqual(t, len(ls), 0, "Размер списка должен быть больше 0.")
}

func TestUpdateAccount(t *testing.T) {

	account := createRandomAccountSimple(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: account.Balance + 10,
	}
	ac, err := testQueries.UpdateAccount(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, ac)
	require.Equal(t, account.ID, ac.ID, "Идентификаторы должны иметь одинакавые значения.")
	require.NotEqual(t, account.Balance, ac.Balance, "Баланс должен иметь разные значения.")
}

func TestDeleteAccount(t *testing.T) {

	account := createRandomAccountSimple(t)
	err := testQueries.DeleteAccount(ctx, account.ID)
	require.NoError(t, err)
	_, errGet := testQueries.GetAccount(ctx, account.ID)
	require.Error(t, errGet)
	require.Equal(t, errGet.Error(), "no rows in result set")
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

func createRandomAccountSimple(t *testing.T) Account {
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

	return account
}
