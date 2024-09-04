package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	t.Skip("Skipping test due to In-complete feature")
	store := NewStore(testDB)

	account1 := createRandomTestAccount(t)
	account2 := createRandomTestAccount(t)

	amount := int64(10)
	n := 5

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			res, err := store.TransferTx(context.Background(), TransferTxParams{
				SenderID:   account1.ID,
				ReceiverID: account2.ID,
				Amount:     amount,
			})

			errs <- err
			results <- res
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		res := <-results
		require.NotEmpty(t, res)

		require.Equal(t, res.Transfer.SenderID, account1.ID)
		require.Equal(t, res.Transfer.ReceiverID, account2.ID)
		require.Equal(t, res.Transfer.Amount, amount)
		_, err = store.GetTransfer(context.Background(), res.Transfer.ID)
		require.NoError(t, err)

		require.Equal(t, res.SenderEntry.AccountID, account2.ID)
		require.Equal(t, res.SenderEntry.Amount, -amount)
		_, err = store.GetEntry(context.Background(), res.SenderEntry.ID)
		require.NoError(t, err)

		require.Equal(t, res.ReceiverEntry.AccountID, account1.ID)
		require.Equal(t, res.ReceiverEntry.Amount, amount)
		_, err = store.GetEntry(context.Background(), res.ReceiverEntry.ID)
		require.NoError(t, err)

		senderAccount := res.SenderAccount
		require.NotEmpty(t, senderAccount)
		require.Equal(t, senderAccount.ID, res.SenderAccount.ID)

		recieverAccount := res.ReceiverAccount
		require.NotEmpty(t, recieverAccount)
		require.Equal(t, senderAccount.ID, res.ReceiverAccount.ID)

		diff1 := res.ReceiverAccount.Balance - res.SenderAccount.Balance
		diff2 := recieverAccount.Balance - senderAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff2 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k < n)
	}

	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, updatedAccount2.Balance, updatedAccount1.Balance+int64(n)*amount)
}
