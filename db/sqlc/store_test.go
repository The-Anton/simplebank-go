package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	account1 := createRandomTestAccount(t)
	account2 := createRandomTestAccount(t)

	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	n := 5

	for i := 0; i<n; i++ {
		go func() {
			res, err := testStore.TransferTx(context.Background(), TransferTxInput{
				SenderID: account1.ID,
				ReceiverID: account2.ID,
				Amount: 0,
			})

			errs <- err
			results <- res
		}()
	}

	for i := 0; i<n; i++ {
		err := <- errs
		require.NoError(t, err)

		result := <- results
		require.NotEmpty(t, result)

		require.Equal(t, result.Transfer.SenderID, account1.ID)
		require.Equal(t, result.Transfer.ReceiverID, account2.ID)
		require.Equal(t, result.Transfer.Amount, amount)
		_, err = testStore.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)

		require.Equal(t, result.SenderEntry.AccountID, account2.ID)
		require.Equal(t, result.SenderEntry.Amount, -amount)
		_, err = testStore.GetEntry(context.Background(), result.SenderEntry.ID)
		require.NoError(t, err)

		require.Equal(t, result.ReceiverEntry.AccountID, account1.ID)
		require.Equal(t, result.ReceiverEntry.Amount, amount)
		_, err = testStore.GetEntry(context.Background(), result.ReceiverEntry.ID)
		require.NoError(t, err)

		senderAccount := result.SenderAccount
		require.NotEmpty(t, senderAccount)
		require.Equal(t, senderAccount.ID, result.SenderAccount.ID)

		recieverAccount := result.ReceiverAccount
		require.NotEmpty(t, recieverAccount)
		require.Equal(t, senderAccount.ID, result.ReceiverAccount.ID)

		diff1 := result.ReceiverAccount.Balance - result.SenderAccount.Balance
		diff2 := recieverAccount.Balance - senderAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff2 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1/amount)
		require.True(t, k >= 1 && k < n)
	}

	updatedAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, updatedAccount2.Balance, updatedAccount1.Balance + int64(n)*amount)
}