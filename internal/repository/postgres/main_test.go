package postgres

import (
	"os"
	"testing"

	"github.com/timothypattikawa/simplebank/internal/config"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	v := config.LoadViper()

	conf := config.NewConfiguration(v)

	db := conf.DBConf.NewDbConn()

	testQueries = New(db)
	os.Exit(m.Run())
}
