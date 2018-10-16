package main

import "fmt"
import "github.com/exhuma/gopgstats/concrete"


func main() {
    fetcher, err := concrete.MakeDefaultFetcher("dbname=exhuma")
    act, err := fetcher.Activity()
    fmt.Printf("Fetcher: %T\n", fetcher)
    fmt.Printf("Error: %T\n", err)
    for idx, a := range(act) {
        fmt.Printf("Row %3d: %-20v %5d %v\n", idx, a.UseName, a.PId, a.State)
    }
}
