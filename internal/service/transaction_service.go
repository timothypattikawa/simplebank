package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"github.com/timothypattikawa/simplebank/internal/dto"
	"github.com/timothypattikawa/simplebank/internal/repository"
	sqlc "github.com/timothypattikawa/simplebank/internal/repository/postgres"
)

type TransactionService interface {
	TransactionTransfer(ctx context.Context, arg dto.TransferTxParams) (dto.TransferTxResult, error)
}

type TransactionServiceImpl struct {
	repository repository.TransactionRepository
	db         *pgxpool.Pool
	v          *viper.Viper
}

func (t TransactionServiceImpl) TransactionTransfer(
	ctx context.Context,
	arg dto.TransferTxParams) (dto.TransferTxResult, error) {
	var result dto.TransferTxResult

	err := t.repository.ExecTx(ctx, func(queries *sqlc.Queries) error {

		transfer, err := queries.CreateTransfer(ctx, sqlc.CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		result.Transfer = transfer

		fromEntry, err := queries.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry = fromEntry

		toEntry, err := queries.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry = toEntry

		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount, result.ToAccount, err =
				addMoney(ctx, queries, arg.FromAccountId, -arg.Amount, arg.ToAccountId, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err =
				addMoney(ctx, queries, arg.ToAccountId, arg.Amount, arg.FromAccountId, -arg.Amount)
		}

		return err
	})

	if err != nil {
		return dto.TransferTxResult{}, err
	}
	return result, err
}

func addMoney(
	ctx context.Context,
	q *sqlc.Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 sqlc.Account, account2 sqlc.Account, err error) {
	account1, err = q.AddAccountBalance(ctx, sqlc.AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, sqlc.AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
}

func NewTransactionService(
	tr repository.TransactionRepository,
	db *pgxpool.Pool,
	v *viper.Viper) TransactionService {
	return &TransactionServiceImpl{
		repository: tr,
		db:         db,
		v:          v,
	}
}
