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
