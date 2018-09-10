package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// driver to specifically connect to postgres
	_ "github.com/lib/pq"
)

// Config is a conf for the postgres database
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLConf
}

// SSLConf Represent the ssl configuration
type SSLConf struct {
	CertPath     string
	KeyPath      string
	RootCertPath string
}

// NewConnPool connects to db and return a connection pool
func NewConnPool(postgresDBConf Config) (*sqlx.DB, error) {
	dsn := "postgres://" +
		postgresDBConf.User + ":" +
		postgresDBConf.Password + "@" +
		postgresDBConf.Host + ":" +
		postgresDBConf.Port + "/" +
		postgresDBConf.DbName

	// Secure connection?
	ssl := "?sslmode=disable"
	if postgresDBConf.SSLConf.CertPath != "" &&
		postgresDBConf.SSLConf.KeyPath != "" &&
		postgresDBConf.SSLConf.RootCertPath != "" {
		ssl = fmt.Sprintf(
			"?sslmode=verify-full&sslcert=%s&sslkey=%s&sslrootcert=%s",
			postgresDBConf.SSLConf.CertPath,
			postgresDBConf.SSLConf.KeyPath,
			postgresDBConf.SSLConf.RootCertPath,
		)
	}
	dsn += ssl

	pool, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("NewConnPool: sqlx.Open %v", err)
	}

	errP := pool.Ping()
	if errP != nil {
		return nil, fmt.Errorf("NewConnPool: pool.Ping %v", errP)
	}

	return pool, nil
}
