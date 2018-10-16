package types

import (
    "github.com/lib/pq"
)


type ActivityRow struct {
    DatId int
    DatName string
    PId int
    UseSysId int
    UseName string
    ApplicationName string
    ClientAddress string
    ClientHostname string
    ClientPort int
    BackendStart pq.NullTime
    XactStart pq.NullTime
    QueryStart pq.NullTime
    StateChange pq.NullTime
    WaitEventType string
    WaitEvent string
    State string
    BackendXid string
    BackendXmin string
    Query string
    BackendType string
}
