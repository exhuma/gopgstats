package gopgstats

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DefaultFetcher struct {
	db *sql.DB
}

// Retturns a real DB connection
func MakeDefaultFetcher(dsn string) (DefaultFetcher, error) {
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", dsn)

	var output DefaultFetcher
	if err != nil {
		output = DefaultFetcher{}
	}
	output = DefaultFetcher{db}
	return output, nil
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

func (fetcher DefaultFetcher) Activity() ([]ActivityRow, error) {
	rows, err := fetcher.db.Query(`
        SELECT
            COALESCE(datid, -1),
            COALESCE(datname, '<unknown>'),
            pid,
            COALESCE(usesysid, -1),
            COALESCE(usename, '<unknown>'),
            COALESCE(application_name, '<unknown>'),
            COALESCE(client_addr, '192.0.2.1')::text,
            COALESCE(client_hostname, '<unknown>'),
            COALESCE(client_port, -1),
            backend_start,
            xact_start,
            query_start,
            state_change,
            COALESCE(wait_event_type, '__none__'),
            COALESCE(wait_event, '__none__'),
            COALESCE(state, 'idle'),
            COALESCE(backend_xid::text, '<unknown>'),
            COALESCE(backend_xmin::text, '<unknown>'),
            COALESCE(query, '<unknown>'),
            backend_type
        FROM pg_stat_activity
    `)
	if err != nil {
		return []ActivityRow{}, err
	}
	defer rows.Close()
	output := []ActivityRow{}
	for rows.Next() {
		var row ActivityRow
		err = rows.Scan(
			&row.DatId,
			&row.DatName,
			&row.PId,
			&row.UseSysId,
			&row.UseName,
			&row.ApplicationName,
			&row.ClientAddress,
			&row.ClientHostname,
			&row.ClientPort,
			&row.BackendStart,
			&row.XactStart,
			&row.QueryStart,
			&row.StateChange,
			&row.WaitEventType,
			&row.WaitEvent,
			&row.State,
			&row.BackendXid,
			&row.BackendXmin,
			&row.Query,
			&row.BackendType)
		if err != nil {
			return []ActivityRow{}, err
		}
		output = append(output, row)
	}
	return output, err
}
