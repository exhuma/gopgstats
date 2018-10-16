package gopgstats

import "testing"

func TestActivity(t *testing.T) {
    fetcher, _ := MakeDefaultFetcher("dbname=exhuma")
    result, _ := fetcher.Activity()
    expected := []ActivityRow{};
    if len(result) != len(expected) {
        t.Errorf("%v != %v", result, expected)
    }
}
