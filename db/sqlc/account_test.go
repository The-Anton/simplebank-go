package db

import (
	"context"
	"simplebank-go/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTestAccount(t *testing.T) Account{
	arg := CreateAccountParams {
		AccountName: util.RandomOwner(6),
		Balance: util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t,account)

	require.Equal(t, arg.AccountName, account.AccountName)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account
}

func TestCreatingAccount(t *testing.T) {
	createRandomTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomTestAccount(t)
	acc, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t,acc)

	require.Equal(t, acc.ID, account.ID)
	require.Equal(t, acc.AccountName, account.AccountName)
	require.Equal(t, acc.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)
}

func TestListAccount(t *testing.T) {
	for i := 0; i<20; i++ {
		createRandomTestAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 10,
		Offset: 2,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t,accounts)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomTestAccount(t)

	arg := UpdateAccountParams {
		ID: account.ID,
		Balance: util.RandomAmount(),
	}

	acc, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t,acc)

	require.Equal(t, acc.ID, account.ID)
	require.Equal(t, acc.AccountName, account.AccountName)
	require.Equal(t, acc.Currency, account.Currency)

	require.Equal(t, acc.Balance, arg.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomTestAccount(t)
	
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
	
	acc, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, acc)
}