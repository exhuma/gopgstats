// Example usage
package main

import (
	"database/sql"
	"fmt"
	"github.com/exhuma/gopgstats"
	_ "github.com/lib/pq"
)

func main() {
	var db *sql.DB
	var err error
	dsn := "host=/var/run/postgresql"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

    // --- Globals
	fmt.Println("--- Locks")
	fetcher := gopgstats.MakeDefaultFetcher(db)
	result, err := fetcher.Locks()
	fmt.Println("Error:", err)
	fmt.Println(result)

	fmt.Println("--- Connections")
	conns, err := fetcher.Connections()
	fmt.Println("Error:", err)
	for _, item := range conns {
		fmt.Println(item)
	}

	fmt.Println("--- QueryAges")
	ages, err := fetcher.QueryAges()
	fmt.Println("Error:", err)
	for _, item := range ages {
		fmt.Println(item)
	}

    // --- Locals
	fmt.Println("--- Disk IO")
	diskio, err := fetcher.DiskIOAll(dsn)
	fmt.Println("Error:", err)
	for _, item := range diskio {
		fmt.Println(item)
	}
}
