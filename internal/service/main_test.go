package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timothypattikawa/simplebank/internal/repository"
	sqlc "github.com/timothypattikawa/simplebank/internal/repository/postgres"
	"os"
	"testing"

	"github.com/timothypattikawa/simplebank/internal/config"
)

var testQueries *sqlc.Queries
var testDb *pgxpool.Pool
var repo repository.TransactionRepository

func TestMain(m *testing.M) {
	v := config.LoadViper()

	conf := config.NewConfiguration(v)

	testDb = conf.DBConf.NewDbConn()

	testQueries = sqlc.New(testDb)

	repo = repository.NewTransactionRepository(testDb, nil)

	os.Exit(m.Run())
}
