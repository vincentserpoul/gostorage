package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

// SQLNamedExecuter is an interface used for CUD
type SQLNamedExecuter interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (
		sql.Result, error,
	)
}

// QueryFilter allows for select filters
type QueryFilter struct {
	FilterSQL   string
	NamedParams map[string]interface{}
}

// SQLNamedQueryer is used to do select queries
type SQLNamedQueryer interface {
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (
		*sqlx.Rows, error,
	)
}

// ErrSQLNotFound happens when no rows were found
var ErrSQLNotFound = errors.New("not found")
