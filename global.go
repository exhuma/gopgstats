// This file contains statistics which are available from any database. As long
// as you are connected to the PostgreSQL server and have the correct
// permissions these statistics are available.
//
// There are also statistics which are only available if you are connected to
// the database from which you want the statistics from. These can be found in
// `local.go`

package gopgstats

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

func (fetcher DefaultFetcher) Connections() ([]ConnectionsRow, error) {
	query := getMatchingQuery(fetcher, ConnectionsQueries[:])
	rows, err := fetcher.db.Query(query)
	defer rows.Close()

	if err != nil {
		return []ConnectionsRow{}, err
	}

	output := []ConnectionsRow{}
	for rows.Next() {
		var row ConnectionsRow
		err = rows.Scan(
			&row.Username,
			&row.Idle,
			&row.IdleInTransaction,
			&row.Unknown,
			&row.QueryActive,
			&row.Waiting)
		if err != nil {
			return []ConnectionsRow{}, err
		}
		output = append(output, row)
	}
	return output, err
}

func (fetcher DefaultFetcher) QueryAges() ([]QueryAgeRow, error) {
	query := getMatchingQuery(fetcher, QueryAgesQueries[:])
	rows, err := fetcher.db.Query(query)
	defer rows.Close()

	if err != nil {
		return []QueryAgeRow{}, err
	}

	output := []QueryAgeRow{}
	for rows.Next() {
		var row QueryAgeRow
		err = rows.Scan(
			&row.DatabaseName,
			&row.QueryAge,
			&row.TransactionAge,
		)
		if err != nil {
			return []QueryAgeRow{}, err
		}
		output = append(output, row)
	}
	return output, err
}

// TODO Transactions
// TODO TempBytes
