package gopgstats

import (
    "testing"
    "fmt"
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
    for _, item := range(result) {
        fmt.Println(item)
    }
    if len(result) == 0 {
        t.Errorf("We did not get any results from the DB")
    }
}
