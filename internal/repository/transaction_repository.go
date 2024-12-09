package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	sqlc "github.com/timothypattikawa/simplebank/internal/repository/postgres"
)

type TransactionRepository interface {
	ExecTx(ctx context.Context, fn func(queries *sqlc.Queries) error) error
}

type TransactionRepositoryImpl struct {
	db *pgxpool.Pool
	v  *viper.Viper
}

func (t TransactionRepositoryImpl) ExecTx(ctx context.Context, fn func(queries *sqlc.Queries) error) error {
	tx, err := t.db.Begin(ctx)
	q := sqlc.New(t.db).WithTx(tx)
	if err != nil {
		return err
	}

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %w, rbErr: %w", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

func NewTransactionRepository(db *pgxpool.Pool, v *viper.Viper) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
		v:  v,
	}
}
