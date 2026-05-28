
package main

import (
    "fmt"
    "time"
)

func worker() {
    fmt.Println("working")
}

func main() {
    go worker()

    time.Sleep(time.Second)
}
