package config

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type connection struct {
	maxIdle           int
	maxLifeTime       time.Duration
	keepAliveInterval time.Duration
	maxOpen           int
}

type DBConf struct {
	cfgName, host, username, password, schema string
	port                                      int
	connection
}

func (conf DBConf) getDBConnection() string {
	connURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.username, conf.password),
		Host:   fmt.Sprintf("%s:%d", conf.host, conf.port),
		Path:   conf.schema,
	}

	q := connURL.Query()
	q.Add("sslmode", "disable")
	// q.Add("TimeZone", "Asia/Jakarta")
	connURL.RawQuery = q.Encode()

	return connURL.String()
}

func (conf DBConf) loadDatabase() *pgxpool.Pool {
	connString := conf.getDBConnection()
	fmt.Println(connString)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	preparePgxConn, err := pgxpool.ParseConfig(connString)

	if err != nil {
		log.Fatalf("%v", err.Error())
	}

	preparePgxConn.MinConns = int32(conf.maxIdle)
	preparePgxConn.MaxConns = int32(conf.maxOpen)
	preparePgxConn.MaxConnLifetime = conf.maxLifeTime
	preparePgxConn.HealthCheckPeriod = conf.keepAliveInterval

	pgxPool, err := pgxpool.NewWithConfig(ctx, preparePgxConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	rows, err := pgxPool.Query(ctx, "select 1")
	if err != nil {
		fmt.Printf("Initial DB fail error {%v} database{%v}", err.Error(), conf.cfgName)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Initial DB fail error {%v} database{%v}", err.Error(), conf.cfgName)
	}

	return pgxPool
}

func (dbConf DBConf) NewDbConn() *pgxpool.Pool {
	return dbConf.loadDatabase()
}

func getConfDBByName(confName string, v *viper.Viper) DBConf {
	newDBConf := DBConf{}

	newDBConf.cfgName = confName
	newDBConf.username = v.GetString(fmt.Sprintf("db.%s.username", confName))
	newDBConf.password = v.GetString(fmt.Sprintf("db.%s.password", confName))
	newDBConf.host = v.GetString(fmt.Sprintf("db.%s.host", confName))
	newDBConf.port = v.GetInt(fmt.Sprintf("db.%s.port", confName))
	newDBConf.schema = v.GetString(fmt.Sprintf("db.%s.schema", confName))

	newDBConf.keepAliveInterval = v.GetDuration(fmt.Sprintf("db.%s.keep-alive-interval", confName))
	newDBConf.maxIdle = v.GetInt(fmt.Sprintf("db.%s.max-idle", confName))
	newDBConf.maxLifeTime = v.GetDuration(fmt.Sprintf("db.%s.max-life-time", confName))
	newDBConf.maxOpen = v.GetInt(fmt.Sprintf("db.%s.max-open", confName))

	return newDBConf
}
