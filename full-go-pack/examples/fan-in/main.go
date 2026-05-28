
package main

import "fmt"

func producer(start int) <-chan int {
    out := make(chan int)

    go func() {
        for i := start; i < start+5; i++ {
            out <- i
        }

        close(out)
    }()

    return out
}

func fanIn(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        for ch1 != nil || ch2 != nil {
            select {
            case v, ok := <-ch1:
                if !ok {
                    ch1 = nil
                    continue
                }
                out <- v
            case v, ok := <-ch2:
                if !ok {
                    ch2 = nil
                    continue
                }
                out <- v
            }
        }

        close(out)
    }()

    return out
}

func main() {
    merged := fanIn(producer(1), producer(100))

    for v := range merged {
        fmt.Println(v)
    }
}
