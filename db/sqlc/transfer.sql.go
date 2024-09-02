// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  sender_id, receiver_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING id, sender_id, receiver_id, amount, currency, created_at
`

type CreateTransferParams struct {
	SenderID   int64 `json:"sender_id"`
	ReceiverID int64 `json:"receiver_id"`
	Amount     int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.SenderID, arg.ReceiverID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.Amount,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, sender_id, receiver_id, amount, currency, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.Amount,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfer = `-- name: ListTransfer :many
SELECT id, sender_id, receiver_id, amount, currency, created_at FROM transfers
WHERE 
    sender_id = $1 OR
    receiver_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransferParams struct {
	SenderID   int64 `json:"sender_id"`
	ReceiverID int64 `json:"receiver_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListTransfer(ctx context.Context, arg ListTransferParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfer,
		arg.SenderID,
		arg.ReceiverID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.SenderID,
			&i.ReceiverID,
			&i.Amount,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
