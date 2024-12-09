package dto

import sqlc "github.com/timothypattikawa/simplebank/internal/repository/postgres"

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id,omitempty"`
	ToAccountId   int64 `json:"to_account_id,omitempty"`
	Amount        int64 `json:"amount,omitempty"`
}

type TransferTxResult struct {
	Transfer    sqlc.Transfer `json:"transfer"`
	FromAccount sqlc.Account  `json:"from_account"`
	ToAccount   sqlc.Account  `json:"to_account"`
	FromEntry   sqlc.Entry    `json:"from_entry"`
	ToEntry     sqlc.Entry    `json:"to_entry"`
}
