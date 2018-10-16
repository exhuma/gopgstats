package stats

import "github.com/exhuma/gopgstats/types"

type StatFetcher interface {
    Activity(...interface{}) ([]types.ActivityRow, error)
}


type DBHandle struct {
    db StatFetcher
}


func (db *DBHandle) Activity() ([]types.ActivityRow, error) {
    // TODO implement
    return []types.ActivityRow{}, nil
}
