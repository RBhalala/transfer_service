package repository

import (
	"context"
	"database/sql"
)

// DBTX is a minimal abstraction over *sql.DB / *sql.Tx.
//
// It lets the same repository code work with “regular” connections
type IDbtx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
