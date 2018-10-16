package main

import "fmt"
import "github.com/exhuma/gopgstats/concrete"


func main() {
    conn, err := concrete.MakeDefaultFetcher("dbname=exhuma")
    act, err := conn.Activity()
    fmt.Println(conn)
    fmt.Println(err)
    fmt.Println(act)
}
