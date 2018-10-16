package main

import "fmt"
import "github.com/exhuma/gopgstats/concrete"


func main() {
    conn, err := concrete.MakeDefaultFetcher("dbname=exhuma")
    act, err := conn.Activity()
    fmt.Println(conn)
    fmt.Println(err)
    for idx, a := range(act) {
        fmt.Printf("Row %03d: %v\n", idx, a)
    }
}
