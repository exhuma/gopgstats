// This file contains statistics which are available only if you are connected
// to the database from which you want the statistics. This means a new
// connection is required per database.
//
// There are also statistics which are always available. These can be found in
// `global.go`

package gopgstats

import "database/sql"

func (fetcher DefaultFetcher) DiskIO(databases []string, dsn string) ([]DiskIOsRow, error) {
	var err error
	output := []DiskIOsRow{}

	for _, dbname := range databases {

		// We need to open a new connection to get access to these stats.
		newDsn := DsnForDatabase(dsn, dbname)
		localDb, err := sql.Open("postgres", newDsn)
		if err != nil {
			return []DiskIOsRow{}, err
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
			return []DiskIOsRow{}, err
		}
		for rows.Next() {
			var row DiskIOsRow
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
				return []DiskIOsRow{}, err
			}
			output = append(output, row)
		}
	}
	return output, err
}

func (fetcher DefaultFetcher) DiskIOAll(dsn string) ([]DiskIOsRow, error) {
	allDbs, err := fetcher.ListDatabases()
	if err != nil {
		return []DiskIOsRow{}, err
	}
	dbs := make([]string, len(allDbs))
	for idx, row := range allDbs {
		dbs[idx] = row.Name
	}
	output, err := fetcher.DiskIO(dbs, dsn)
	return output, err
}

func (fetcher DefaultFetcher) IndexIO(databases []string, dsn string) ([]IndexIOsRow, error) {
	var err error
	output := []IndexIOsRow{}

	for _, dbname := range databases {

		// We need to open a new connection to get access to these stats.
		newDsn := DsnForDatabase(dsn, dbname)
		localDb, err := sql.Open("postgres", newDsn)
		if err != nil {
			return []IndexIOsRow{}, err
		}
		defer localDb.Close()

		// TODO It is ugly that we use "fetcher" to determine the db version
		// but run the query on "localDb". It would be better to have a method
		// in "fetcher" which executes a query on a localised connection, so
		// instead of having "fetcher.db.Query", it would be better to split it
		// into "fetcher.Query" and "fetcher.LocalQuery".
		query := getMatchingQuery(fetcher, IndexIOQueries[:])
		rows, err := localDb.Query(query)
		defer rows.Close()
		if err != nil {
			return []IndexIOsRow{}, err
		}
		for rows.Next() {
			var row IndexIOsRow
			row.DatabaseName = dbname
			err = rows.Scan(
				&row.IndexBlocksRead,
				&row.IndexBlocksHit,
			)
			if err != nil {
				return []IndexIOsRow{}, err
			}
			output = append(output, row)
		}
	}
	return output, err
}

func (fetcher DefaultFetcher) IndexIOAll(dsn string) ([]IndexIOsRow, error) {
	allDbs, err := fetcher.ListDatabases()
	if err != nil {
		return []IndexIOsRow{}, err
	}
	dbs := make([]string, len(allDbs))
	for idx, row := range allDbs {
		dbs[idx] = row.Name
	}
	output, err := fetcher.IndexIO(dbs, dsn)
	return output, err
}

func (fetcher DefaultFetcher) SequencesIO(databases []string, dsn string) ([]SequencesIOsRow, error) {
	var err error
	output := []SequencesIOsRow{}

	for _, dbname := range databases {

		// We need to open a new connection to get access to these stats.
		newDsn := DsnForDatabase(dsn, dbname)
		localDb, err := sql.Open("postgres", newDsn)
		if err != nil {
			return []SequencesIOsRow{}, err
		}
		defer localDb.Close()

		// TODO It is ugly that we use "fetcher" to determine the db version
		// but run the query on "localDb". It would be better to have a method
		// in "fetcher" which executes a query on a localised connection, so
		// instead of having "fetcher.db.Query", it would be better to split it
		// into "fetcher.Query" and "fetcher.LocalQuery".
		query := getMatchingQuery(fetcher, SequencesIOQueries[:])
		rows, err := localDb.Query(query)
		defer rows.Close()
		if err != nil {
			return []SequencesIOsRow{}, err
		}
		for rows.Next() {
			var row SequencesIOsRow
			row.DatabaseName = dbname
			err = rows.Scan(
				&row.BlocksRead,
				&row.BlocksHit,
			)
			if err != nil {
				return []SequencesIOsRow{}, err
			}
			output = append(output, row)
		}
	}
	return output, err
}

func (fetcher DefaultFetcher) SequencesIOAll(dsn string) ([]SequencesIOsRow, error) {
	allDbs, err := fetcher.ListDatabases()
	if err != nil {
		return []SequencesIOsRow{}, err
	}
	dbs := make([]string, len(allDbs))
	for idx, row := range allDbs {
		dbs[idx] = row.Name
	}
	output, err := fetcher.SequencesIO(dbs, dsn)
	return output, err
}

func (fetcher DefaultFetcher) ScanTypes(databases []string, dsn string) ([]ScanTypesRow, error) {
	var err error
	output := []ScanTypesRow{}

	for _, dbname := range databases {

		// We need to open a new connection to get access to these stats.
		newDsn := DsnForDatabase(dsn, dbname)
		localDb, err := sql.Open("postgres", newDsn)
		if err != nil {
			return []ScanTypesRow{}, err
		}
		defer localDb.Close()

		// TODO It is ugly that we use "fetcher" to determine the db version
		// but run the query on "localDb". It would be better to have a method
		// in "fetcher" which executes a query on a localised connection, so
		// instead of having "fetcher.db.Query", it would be better to split it
		// into "fetcher.Query" and "fetcher.LocalQuery".
		query := getMatchingQuery(fetcher, ScanTypesQueries[:])
		rows, err := localDb.Query(query)
		defer rows.Close()
		if err != nil {
			return []ScanTypesRow{}, err
		}
		for rows.Next() {
			var row ScanTypesRow
			row.DatabaseName = dbname
			err = rows.Scan(
				&row.IndexScans,
				&row.SequentialScans,
			)
			if err != nil {
				return []ScanTypesRow{}, err
			}
			output = append(output, row)
		}
	}
	return output, err
}

func (fetcher DefaultFetcher) ScanTypesAll(dsn string) ([]ScanTypesRow, error) {
	allDbs, err := fetcher.ListDatabases()
	if err != nil {
		return []ScanTypesRow{}, err
	}
	dbs := make([]string, len(allDbs))
	for idx, row := range allDbs {
		dbs[idx] = row.Name
	}
	output, err := fetcher.ScanTypes(dbs, dsn)
	return output, err
}
