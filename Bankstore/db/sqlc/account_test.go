/*
Pfttern: <filename>_test.go
*/
package db

import (
	"Bankstore/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
func TestGetAccount(t *testing.T) {
	account := createRandomAccountSimple(t)
	a, err := testQueries.GetAccount(context.Background(), account.ID)
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
	ls, err := testQueries.ListAccounts(context.Background(), arg)
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
	ac, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ac)
	require.Equal(t, account.ID, ac.ID, "Идентификаторы должны иметь одинакавые значения.")
	require.NotEqual(t, account.Balance, ac.Balance, "Баланс должен иметь разные значения.")
}

func TestDeleteAccount(t *testing.T) {

	account := createRandomAccountSimple(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
	_, errGet := testQueries.GetAccount(context.Background(), account.ID)
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
	account, err := testQueries.CreateAccount(context.Background(), arg)

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
	account, err := testQueries.CreateAccount(context.Background(), arg)

	// Две проверки на результат работы CreateAccount
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}
