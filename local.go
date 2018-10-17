// This file contains statistics which are available only if you are connected
// to the database from which you want the statistics. This means a new
// connection is required per database.
//
// There are also statistics which are always available. These can be found in
// `global.go`

package gopgstats

import "database/sql"

func (fetcher DefaultFetcher) DiskIO(databases []string, dsn string) ([]DiskIORow, error) {
	var err error
	output := []DiskIORow{}

	for _, dbname := range databases {

		// We need to open a new connection to get access to these stats.
		newDsn := DsnForDatabase(dsn, dbname)
		localDb, err := sql.Open("postgres", newDsn)
		if err != nil {
			return []DiskIORow{}, err
		}
		defer localDb.Close()

		// TODO It is ugly that we use "fetcher" to determine the db version
		// but run the query on "localDb". It would be better to have a method
		// in "fetcher" which executes a query on a localised connection, so
		// instead of having "fetcher.db.Query", it would be better to split it
		// into "fetcher.Query" and "fetcher.LocalQuery".
		query := getMatchingQuery(fetcher, DiskIOQueries[:])
		rows, err := localDb.Query(query)
		defer rows.Close()
		if err != nil {
			return []DiskIORow{}, err
		}
		for rows.Next() {
			var row DiskIORow
			row.DatabaseName = dbname
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
	}
	return output, err
}

func (fetcher DefaultFetcher) DiskIOAll(dsn string) ([]DiskIORow, error) {
	allDbs, err := fetcher.ListDatabases()
	if err != nil {
		return []DiskIORow{}, err
	}
	dbs := make([]string, len(allDbs))
	for idx, row := range allDbs {
		dbs[idx] = row.Name
	}
	output, err := fetcher.DiskIO(dbs, dsn)
	return output, err
}
