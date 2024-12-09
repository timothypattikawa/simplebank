package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/timothypattikawa/simplebank/pkg/utils"
)

func CreateRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	result, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.Owner, args.Owner)
	require.Equal(t, result.Balance, args.Balance)
	require.Equal(t, result.Currency, args.Currency)

	require.NotZero(t, result.ID)

	return result
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	result, err := testQueries.GetAccountById(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.Owner, account1.Owner)
	require.Equal(t, result.Balance, account1.Balance)
	require.Equal(t, result.Currency, account1.Currency)

	require.NotZero(t, result.ID)
	require.Equal(t, account1.ID, result.ID)
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{ID: account1.ID, Balance: utils.RandomMoney()}
	result, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.Balance, result.Balance)

	require.NotZero(t, result.ID)
	require.Equal(t, arg.ID, result.ID)
}
