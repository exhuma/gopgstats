package gopgstats

type VersionedQuery struct {
    minVersion int
    query string
}


func GetFirstMatch(collection []VersionedQuery, minVersion int) (string, bool) {
    lastVisited := ""
    for _, item := range(collection) {
        if item.minVersion == minVersion {
            return item.query, true
        }
        if item.minVersion > minVersion {
            return lastVisited, lastVisited != ""
        }
        lastVisited = item.query
    }
    return lastVisited, lastVisited != ""
}
