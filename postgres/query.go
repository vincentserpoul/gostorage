package postgres

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

// SQLPrepareNamed is THE most useful interface
// as postgres requires two phase to insert, with a prepare stmt
type SQLPrepareNamed interface {
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}

// SQLNamedQueryer is used to do select queries
type SQLNamedQueryer interface {
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (
		*sqlx.Rows, error,
	)
}

// QueryFilter allows for select filters
type QueryFilter struct {
	FilterSQL   string
	NamedParams map[string]interface{}
}

// ErrSQLNotFound happens when no rows were found
var ErrSQLNotFound = errors.New("not found")
