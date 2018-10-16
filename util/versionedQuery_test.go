package util

import "testing"


func TestGetFirstMatch(t *testing.T) {
    var collection [3]VersionedQuery;
    collection[0] = VersionedQuery{100, "100"}
    collection[1] = VersionedQuery{200, "200"}
    collection[2] = VersionedQuery{300, "300"}


    tables := []struct {
        key int
        query string
        valid bool
    }{
        {50, "", false},
        {100, "100", true},
        {120, "100", true},
        {200, "200", true},
        {250, "200", true},
        {300, "300", true},
        {999, "300", true},
    }

    for _, table := range(tables) {
        got, valid := GetFirstMatch(collection[:], table.key)
        want := table.query
        want_valid := table.valid
        if got != want || valid != want_valid {
            t.Errorf("GetFirstMatch(..., %v) == (%q, %v), want (%q, %v)",
                     table.key, got, valid, want, want_valid)
        }
    }

}
