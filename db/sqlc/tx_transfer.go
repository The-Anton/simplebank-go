package db

import (
	"context"
	"fmt"
)

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	SenderID   int64 `json:"sender_id"`
	ReceiverID int64 `json:"reciever_id"`
	Amount     int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer        Transfer `json:"transfers"`
	SenderAccount   Account  `json:"sender_account"`
	ReceiverAccount Account  `json:"reciever_account"`
	SenderEntry     Entry    `json:"sender_entry"`
	ReceiverEntry   Entry    `json:"reciever_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			SenderID:   arg.SenderID,
			ReceiverID: arg.ReceiverID,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}

		result.SenderEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ReceiverID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ReceiverEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.SenderID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// result.SenderAccount, err = q.GetAccountForUpdate(ctx, result.Transfer.SenderID)
		// if err != nil {
		// 	return err
		// }

		result.SenderAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Amount: -arg.Amount,
			ID:     arg.SenderID,
		})
		if err != nil {
			return err
		}

		// result.ReceiverAccount, err = q.GetAccountForUpdate(ctx, result.Transfer.ReceiverID)
		// if err != nil {
		// 	return err
		// }

		result.ReceiverAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Amount: arg.Amount,
			ID:     arg.ReceiverID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
