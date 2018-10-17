// This file contains helpers for DB mocking and unit tests.
// While there is no strong unit-testing in use yet, the function
// MakeDefaultFetcher is considered the main entry-point of this package.

package gopgstats

import (
	"database/sql"
)

type DefaultFetcher struct {
	db *sql.DB
}

// Retturns a real DB connection
func MakeDefaultFetcher(db *sql.DB) DefaultFetcher {
	var output DefaultFetcher
	output = DefaultFetcher{db}
	return output
}
