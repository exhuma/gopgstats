package gopgstats

var DiskSizeQueries = [1]VersionedQuery{
	VersionedQuery{0, `SELECT datname, pg_database_size(datname)
                       FROM pg_database WHERE datistemplate=false`}}

var LocksQueries = [1]VersionedQuery{
	VersionedQuery{0, `SELECT
        COALESCE(db.datname, '<unknown>'),
        COALESCE(LOWER(mode), '<unknown>'),
        COALESCE(locktype, '<unknown>'),
        COALESCE(granted, false),
        COUNT(mode)
    FROM pg_database db
    FULL OUTER JOIN pg_locks lck ON (db.oid=lck.database)
    GROUP BY db.datname, mode, locktype, granted`}}
