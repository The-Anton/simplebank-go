package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestCreatingAccount(t *testing.T) {
	arg := CreateAccountParams {
		AccountName: "Shekhar",
		Balance: 100,
		Currency: "INR",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t,account)

	require.Equal(t, arg.AccountName, account.AccountName)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)


}