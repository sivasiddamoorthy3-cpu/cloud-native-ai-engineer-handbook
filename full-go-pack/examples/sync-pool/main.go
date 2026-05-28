
package main

import (
    "bytes"
    "fmt"
    "sync"
)

func main() {
    pool := sync.Pool{
        New: func() interface{} {
            return new(bytes.Buffer)
        },
    }

    buf := pool.Get().(*bytes.Buffer)

    buf.WriteString("hello")

    fmt.Println(buf.String())

    buf.Reset()

    pool.Put(buf)
}
