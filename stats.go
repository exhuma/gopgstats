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

func getMatchingQuery(fetcher DefaultFetcher, queries []VersionedQuery) string {
	rows, err := fetcher.db.Query("select current_setting('server_version_num')")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	var version int
	rows.Next()
	rows.Scan(&version)
	query, valid := GetFirstMatch(queries, version)
	if !valid {
		panic("Unable to get a matching query for this db-version!")
	}
	return query
}

func (fetcher DefaultFetcher) DiskSize() ([]DiskSizeRow, error) {
	query := getMatchingQuery(fetcher, DiskSizeQueries[:])
	rows, err := fetcher.db.Query(query)
	defer rows.Close()

	if err != nil {
		return []DiskSizeRow{}, err
	}

	output := []DiskSizeRow{}
	for rows.Next() {
		var row DiskSizeRow
		err = rows.Scan(&row.DatabaseName, &row.Size)
		if err != nil {
			return []DiskSizeRow{}, err
		}
		output = append(output, row)
	}
	return output, err
}

func (fetcher DefaultFetcher) DiskIO() ([]DiskIORow, error) {
	query := getMatchingQuery(fetcher, DiskIOQueries[:])
	rows, err := fetcher.db.Query(query)
	defer rows.Close()

	if err != nil {
		return []DiskIORow{}, err
	}

	output := []DiskIORow{}
	for rows.Next() {
		var row DiskIORow
		err = rows.Scan(
			&row.HeapBlocksRead,
			&row.HeapBlocksHit,
			&row.IndexBlocksRead,
			&row.IndexBlocksHit,
			&row.ToastBlocksRead,
			&row.ToastBlocksHit,
			&row.ToastIndexBlocksRead,
			&row.ToastIndexBlocksHit)
		if err != nil {
			return []DiskIORow{}, err
		}
		output = append(output, row)
	}
	return output, err
}

func (fetcher DefaultFetcher) Locks() ([]LocksRow, error) {
	query := getMatchingQuery(fetcher, LocksQueries[:])
	rows, err := fetcher.db.Query(query)
	defer rows.Close()

	if err != nil {
		return []LocksRow{}, err
	}

	output := []LocksRow{}
	for rows.Next() {
		var row LocksRow
		err = rows.Scan(
			&row.DatabaseName,
			&row.Mode,
			&row.Type,
			&row.Granted,
			&row.Count)
		if err != nil {
			return []LocksRow{}, err
		}
		output = append(output, row)
	}
	return output, err
}
