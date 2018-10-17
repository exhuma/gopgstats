// This file contains the needed queries for the database. It is in very close
// relation to `types.go` which defines a struct for each row related to these
// queries. If the structure of a query changes, the corresponding struct must
// also be adapted in `types.go`
//
// In order to accomodate changes in tables structures across PostgreSQL
// versions ,these queries are defined with the help of the "VersionedQuery"
// type. This takes two values. The Minimum version since when the query was
// available and the actual query. Each query in a group *must* have the same
// structure/schema!
//
// The version comparison is done using the `server_version_num` value from
// PostgreSQL, which is a simple numerical value.
//
// A query which is the same across all versions can simply define a minimal
// version of `0`.
package gopgstats

// Retrieves the appropriate query according to the PostgreSQL server version
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

// --- "Global" Query definitions --------------------------------------------

var ListDBQueries = [1]VersionedQuery{
	VersionedQuery{0, `SELECT datname
                       FROM pg_database
                       WHERE datistemplate=false`}}

var DiskSizeQueries = [1]VersionedQuery{
	VersionedQuery{0, `SELECT datname, pg_database_size(datname)
                       FROM pg_database WHERE datistemplate=false`}}

var LocksQueries = [1]VersionedQuery{
	VersionedQuery{0, `SELECT
        COALESCE(db.datname, '<unknown>'),
        COALESCE(LOWER(mode), '<unknown>'),
        COALESCE(locktype, '<unknown>'),
        COALESCE(granted, true),
        COUNT(mode)
    FROM pg_database db
    FULL OUTER JOIN pg_locks lck ON (db.oid=lck.database)
    GROUP BY db.datname, mode, locktype, granted`}}

var QueryAgesQueries = [2]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            datname,
            COALESCE(MAX(extract(EPOCH FROM NOW() - query_start)), 0),
            COALESCE(MAX(extract(EPOCH FROM NOW() - xact_start)), 0)
        FROM pg_stat_activity
        WHERE current_query NOT LIKE '<IDLE%'
        GROUP BY datname`},
	VersionedQuery{90200, `
        SELECT
            datname,
            COALESCE(MAX(extract(EPOCH FROM NOW() - query_start)), 0),
            COALESCE(MAX(extract(EPOCH FROM NOW() - xact_start)), 0)
        FROM pg_stat_activity
        WHERE state NOT LIKE '%idle%'
        GROUP BY datname`}}

var TransactionsQuery = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            datname,
            pg_stat_get_db_xact_commit(oid),
            pg_stat_get_db_xact_rollback(oid)
        FROM pg_database`}}

var TempBytesQueries = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            datname,
            temp_bytes
        FROM
            pg_stat_database`}}

var ConnectionsQueries = [3]VersionedQuery{
	VersionedQuery{0, `
        WITH users AS (SELECT usename FROM pg_user),
        conntype AS (SELECT u.usename,
            act.pid,
            act.waiting,
            current_query
            FROM users u
            LEFT JOIN pg_stat_activity act USING (usename))
        SELECT
            usename,
            COUNT(CASE WHEN current_query='<IDLE>'
                THEN 1 END) AS idle,
            COUNT(CASE WHEN current_query='<IDLE> in transaction'
                THEN 1 END) AS idle_tx,
            COUNT(CASE WHEN current_query='<insufficient privilege>'
                THEN 1 END) AS unknown,
            COUNT(CASE WHEN current_query NOT IN (
                '<IDLE>',
                '<IDLE> in transaction',
                '<insufficient privilege>')
                THEN 1 END) AS query_running,
            COUNT(CASE WHEN waiting THEN 1 END) AS waiting
        FROM conntype
        WHERE COALESCE(conntype.pid, 0) <> pg_backend_pid()
        GROUP BY usename
        ORDER BY usename`},
	VersionedQuery{90200, `
        WITH users AS (SELECT usename FROM pg_user),
        conntype AS (SELECT u.usename,
            act.pid,
            act.waiting,
            state,
            query
            FROM users u
            LEFT JOIN pg_stat_activity act USING (usename))
        SELECT
            usename,
            COUNT(CASE WHEN state = 'idle'
                THEN 1 END) AS idle,
            COUNT(CASE WHEN state like 'idle in transaction%'
                THEN 1 END) AS idle_tx,
            COUNT(CASE WHEN state NOT IN (
                'idle',
                'idle in transaction',
                'idle in transaction (aborted)',
                'active')
                THEN 1 END) AS unknown,
            COUNT(CASE WHEN state = 'active'
                THEN 1 END) AS query_running,
            COUNT(CASE WHEN waiting THEN 1 END) AS waiting
        FROM conntype
        WHERE COALESCE(conntype.pid, 0) <> pg_backend_pid()
        GROUP BY usename
        ORDER BY usename`},
	VersionedQuery{100000, `
        WITH users AS (SELECT usename FROM pg_user),
        conntype AS (SELECT u.usename,
            act.pid,
            act.wait_event_type IS NOT NULL as waiting,
            state,
            query
            FROM users u
            LEFT JOIN pg_stat_activity act USING (usename))
        SELECT
            usename,
            COUNT(CASE WHEN state = 'idle'
                THEN 1 END) AS idle,
            COUNT(CASE WHEN state like 'idle in transaction%'
                THEN 1 END) AS idle_tx,
            COUNT(CASE WHEN state NOT IN (
                'idle',
                'idle in transaction',
                'idle in transaction (aborted)',
                'active')
                THEN 1 END) AS unknown,
            COUNT(CASE WHEN state = 'active'
                THEN 1 END) AS query_running,
            COUNT(CASE WHEN waiting THEN 1 END) AS waiting
        FROM conntype
        WHERE COALESCE(conntype.pid, 0) <> pg_backend_pid()
        GROUP BY usename
        ORDER BY usename`}}

// --- "Local" Query definitions ---------------------------------------------

var DiskIOQueries = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            COALESCE(SUM(heap_blks_read), 0) AS heap_blks_read,
            COALESCE(SUM(heap_blks_hit), 0) AS heap_blks_hit,
            COALESCE(SUM(idx_blks_read), 0) AS idx_blks_read,
            COALESCE(SUM(idx_blks_hit), 0) AS idx_blks_hit,
            COALESCE(SUM(toast_blks_read), 0) AS toast_blks_read,
            COALESCE(SUM(toast_blks_hit), 0) AS toast_blks_hit,
            COALESCE(SUM(tidx_blks_read), 0) AS tidx_blks_read,
            COALESCE(SUM(tidx_blks_hit), 0) AS tidx_blks_hit
        FROM pg_statio_user_tables;`}}

var IndexIOQueries = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            SUM(idx_blks_read) AS idx_blks_read,
            SUM(idx_blks_hit) AS idx_blks_hit
        FROM pg_statio_user_indexes;`}}

var SequencesIOQueries = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            SUM(blks_read) AS blks_read,
            SUM(blks_hit) AS blks_hit
        FROM pg_statio_user_sequences`}}

var ScanTypesQueries = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            SUM(idx_scan) AS idx_scan,
            SUM(seq_scan) AS seq_scan
        FROM pg_stat_user_tables`}}

var RowAccessQueries = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            SUM(n_tup_ins) AS n_tup_ins,
            SUM(n_tup_upd) AS n_tup_upd,
            SUM(n_tup_del) AS n_tup_del,
            SUM(n_tup_hot_upd) AS n_tup_hot_upd
        FROM pg_stat_user_tables`}}

var SizesQueries = [1]VersionedQuery{
	VersionedQuery{0, `
        SELECT
            SUM(pg_relation_size(oid, 'main')) AS main_size,
            SUM(pg_relation_size(oid, 'vm')) AS vm_size,
            SUM(pg_relation_size(oid, 'fsm')) AS fsm_size,
            SUM(
                CASE reltoastrelid
                WHEN 0 THEN 0
                ELSE pg_total_relation_size(reltoastrelid)
                END
            ) AS toast_size,
            SUM(pg_indexes_size(oid)) AS indexes_size,
            pg_database_size(current_database()) AS database_size
            FROM pg_class
            WHERE relkind not in ('t', 'i')
            AND NOT relisshared`}}
