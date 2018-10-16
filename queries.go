package gopgstats


var DiskSizeQueries = [1]VersionedQuery{
    VersionedQuery{0, `SELECT datname, pg_database_size(datname)
                       FROM pg_database WHERE datistemplate=false`}}


var LocksQueries = [1]VersionedQuery{
    VersionedQuery{0, `SELECT
        db.datname,
        LOWER(mode),
        locktype,
        granted,
        COUNT(mode)
    FROM pg_database db
    FULL OUTER JOIN pg_locks lck ON (db.oid=lck.database)
    GROUP BY db.datname, mode, locktype, granted`}}
