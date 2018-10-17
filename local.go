// This file contains statistics which are available only if you are connected
// to the database from which you want the statistics. This means a new
// connection is required per database.
//
// There are also statistics which are always available. These can be found in
// `global.go`

package gopgstats

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
