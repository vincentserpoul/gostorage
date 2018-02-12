package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// driver to specifically connect to mysql
	_ "github.com/go-sql-driver/mysql"
)

// Config is a conf for the mysql database
type Config struct {
	Protocol string
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

// NewConnPool connects to db and return a connection pool
func NewConnPool(mysqlDBConf Config) (*sqlx.DB, error) {
	dsn := mysqlDBConf.User + ":" +
		mysqlDBConf.Password + "@" +
		mysqlDBConf.Protocol + "(" +
		mysqlDBConf.Host + ":" +
		mysqlDBConf.Port + ")/" +
		mysqlDBConf.DbName + "?parseTime=true&multiStatements=true"

	pool, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("NewConnPool: sqlx.Open %v", err)
	}

	errP := pool.Ping()
	if errP != nil {
		return nil, fmt.Errorf("NewConnPool: pool.Ping %v", errP)
	}

	return pool, nil
}
