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

	pool, err := sqlx.Open("postgres", getDSN(postgresDBConf))
	if err != nil {
		return nil, fmt.Errorf("NewConnPool: sqlx.Open %v", err)
	}

	errP := pool.Ping()
	if errP != nil {
		return nil, fmt.Errorf("NewConnPool: pool.Ping %v", errP)
	}

	return pool, nil
}

func getDSN(c Config) string {
	dsn := "postgres://" +
		c.User + ":" +
		c.Password + "@" +
		c.Host + ":" +
		c.Port + "/" +
		c.DbName

	// Secure connection?
	ssl := "?sslmode=disable"
	if c.SSLConf.CertPath != "" &&
		c.SSLConf.KeyPath != "" &&
		c.SSLConf.RootCertPath != "" {
		ssl = fmt.Sprintf(
			"?sslmode=verify-full&sslcert=%s&sslkey=%s&sslrootcert=%s",
			c.SSLConf.CertPath,
			c.SSLConf.KeyPath,
			c.SSLConf.RootCertPath,
		)
	}

	return dsn + ssl
}
