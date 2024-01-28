package dbs

import (
	"context"
	"github.com/stretchr/testify/require"
	"sim_bank/util"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.Owner, arg.Owner)
	require.Equal(t, acc.Balance, arg.Balance)

	require.NotZero(t, acc.ID)
	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.NoError(t, err)
	require.Equal(t, acc.ID, acc2.ID)
	require.Equal(t, acc.Owner, acc2.Owner)
}
