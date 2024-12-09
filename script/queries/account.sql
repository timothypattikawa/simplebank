-- name: CreateAccount :one
INSERT INTO accounts ("owner", balance, currency)
VALUES($1, $2, $3) RETURNING *;

-- name: GetAccountById :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: ListOfAccount :many
SELECT * FROM accounts ORDER BY id LIMIT $1;

-- name: UpdateAccount :one
UPDATE accounts SET balance = $2 WHERE id = $1 RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
    RETURNING *;