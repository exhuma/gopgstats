package gopgstats


type StatFetcher interface {
    Activity(...interface{}) ([]ActivityRow, error)
}


type DBHandle struct {
    db StatFetcher
}


func (db *DBHandle) Activity() ([]ActivityRow, error) {
    // TODO implement
    return []ActivityRow{}, nil
}
