package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	conn, _ := sql.Open(dbDriver, dbSource)
	testQueries = New(conn)

	arg := CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {

	conn, _ := sql.Open(dbDriver, dbSource)
	testQueries = New(conn)

	arg := CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)

	fmt.Printf("account: %v\n", account)
	fmt.Printf("account2: %v\n", account2)

}

func TestDeleteAccount(t *testing.T) {

	conn, _ := sql.Open(dbDriver, dbSource)
	testQueries = New(conn)

	arg := CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	err = testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestDeleteAccountById(t *testing.T) {

	conn, _ := sql.Open(dbDriver, dbSource)
	testQueries = New(conn)

	for idx := 2; idx < 10; idx++ {
		err := testQueries.DeleteAccount(context.Background(), int64(idx))
		require.NoError(t, err)
	}

}

func TestListAccounts(t *testing.T) {

	conn, _ := sql.Open(dbDriver, dbSource)
	testQueries = New(conn)

	for i := 0; i < 3; i++ {
		arg := CreateAccountParams{
			Owner:    "Tom",
			Balance:  100,
			Currency: "USD",
		}
		_, err := testQueries.CreateAccount(context.Background(), arg)
		require.NoError(t, err)
	}

	for i := 0; i < 3; i++ {
		arg := CreateAccountParams{
			Owner:    "SAM",
			Balance:  100,
			Currency: "USD",
		}
		_, err := testQueries.CreateAccount(context.Background(), arg)
		require.NoError(t, err)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	result, _ := testQueries.ListAccount(context.Background(), arg)

	for _, account := range result {
		fmt.Printf("account: %v\n", account)
	}

}
