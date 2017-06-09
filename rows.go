package awql

import "errors"

var (
	ErrNoRows = errors.New("no more rows")
)

// Iterator for query results.
type Rows interface {
	Next() bool
	Scan() (map[string]interface{}, error)
	Close()
}
