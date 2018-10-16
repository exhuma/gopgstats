package gopgstats

import "testing"

func TestActivity(t *testing.T) {
    fetcher, _ := MakeDefaultFetcher("dbname=exhuma")
    result, _ := fetcher.Activity()
    if len(result) == 0 {
        t.Errorf("We did not get any results from the DB")
    }
}
