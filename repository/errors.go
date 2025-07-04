package repository

import "errors"

// ErrNoData is returned when a query returns no rows.
var ErrNoData = errors.New("No Data Found")
var ErrDupData = errors.New("Duplicate Record")
