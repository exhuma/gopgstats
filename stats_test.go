package gopgstats

import (
	"fmt"
	"testing"
)

func TestActivity(t *testing.T) {
	fetcher, _ := MakeDefaultFetcher("dbname=exhuma")
	result, _ := fetcher.Activity()
	if len(result) == 0 {
		t.Errorf("We did not get any results from the DB")
	}
}

func TestDisksize(t *testing.T) {
	fetcher, _ := MakeDefaultFetcher("dbname=exhuma")
	result, _ := fetcher.DiskSize()
	for _, item := range result {
		fmt.Println("disksize: ", item)
	}
}

func TestLocks(t *testing.T) {
	fetcher, _ := MakeDefaultFetcher("dbname=exhuma")
	result, _ := fetcher.Locks()
	for _, item := range result {
		fmt.Println("locks: ", item)
	}
}
