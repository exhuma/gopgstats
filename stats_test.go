package gopgstats

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestDisksize(t *testing.T) {
	db, _ := sql.Open("postgres", "dbname=exhuma")
	fetcher := MakeDefaultFetcher(db)
	result, _ := fetcher.DiskSize()
	for _, item := range result {
		fmt.Println("disksize: ", item)
	}
}

func TestLocks(t *testing.T) {
	db, _ := sql.Open("postgres", "dbname=exhuma")
	fetcher := MakeDefaultFetcher(db)
	result, _ := fetcher.Locks()
	for _, item := range result {
		fmt.Println("locks: ", item)
	}
}
