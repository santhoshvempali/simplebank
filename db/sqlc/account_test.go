package db

import (
	"context"
	"log"
	"testing"

	"github.com/santhoshvempali/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account

}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1, account2)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: util.RandomMoney(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
}
