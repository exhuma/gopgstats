package gopgstats


import (
    _ "github.com/lib/pq"
    "database/sql"
)


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


type DefaultFetcher struct {
    db *sql.DB
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
    output := []ActivityRow{};
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
