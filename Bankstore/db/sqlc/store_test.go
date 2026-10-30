package db

import (
	"context"
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println("Before:", account1.Balance, account2.Balance)
	n:=5
	amount := int64(10)
	errs:=make(chan error)
	results:=make(chan TransferTxResult)
	// Run n concurrent transfer transactions
	for range 5{
		go func ()  {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs<-err
			results<-result
		}()
	}
	// chack results
	for range n{
		// chack errors
		err:=<-errs
		require.NoError(t, err)

		//chack results
		result:=<-results
		transfer:=result.Transfer
		require.NotEmpty(t,transfer)
		require.Equal(t,account1.ID,transfer.FromAccountID)
		require.Equal(t,account2.ID,transfer.ToAccountID)
		require.Equal(t,amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		//Additionsl chack
		_, err=store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)
	}
	
}
