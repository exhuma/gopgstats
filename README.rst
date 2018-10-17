PostgreSQL statistics for Go
============================


This package contains helper functions to retrieve runtime statistcs of a
PostgreSQL DB server. The aim is to decouple the DB queries from Go
applications.

Statistics will always return as a list of structs where each statistic has its
own struct. The decision to do this instead of returning a list of
``interface{}`` items was that this keeps type-safety and removes guesswork.

As of 2018-10-17, this package contains a couple of statistics. More will come
soon.


DSN & URLs
----------

Please take note that URLs are **not** yet supported as database DSNs. Certain
statistics are only available on the databse of the currently active
connection. This means that in order to fetch statistics from all databases, a
new connection must be established per DB. ``gopgstats`` will reuse the DSN
which was used to create the ``Fetcher`` object and modify it as needed. This
modification is currently only coded for normal connection strings. Not for
URLs!


Usage
-----

::

    import (
        "database/sql"
        "github.com/exhuma/gopgstats"
        _ "github.com/lib/pq"
    )

    // Use a normal PostgreSQL DSN syntax here. Note that URLs are currently
    // NOT supported!
    dsn := "host=dbhost username=jdoe password=s3cr37"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        panic(err)
    }

    fmt.Println("--- Locks")
    fetcher := gopgstats.MakeDefaultFetcher(db)
    result, err := fetcher.Locks()
    fmt.Println(err)
    fmt.Println(result)

    fmt.Println("--- Disk IO")
    diskio, err := fetcher.DiskIOAll(dsn)
    fmt.Println(err)
    fmt.Println(diskio)
