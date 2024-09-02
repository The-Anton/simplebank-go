-- name: CreateTransfer :one
INSERT INTO transfers (
  sender_id, receiver_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfer :many
SELECT * FROM transfers
WHERE 
    sender_id = $1 OR
    receiver_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

