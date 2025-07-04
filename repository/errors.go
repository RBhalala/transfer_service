package repository

import "errors"

// ErrNoData is returned when a query returns no rows.
var ErrNoData = errors.New("no data found")
