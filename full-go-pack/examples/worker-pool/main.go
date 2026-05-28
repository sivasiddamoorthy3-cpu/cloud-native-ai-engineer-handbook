
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    for job := range jobs {
        fmt.Printf("worker %d processing %d\n", id, job)
    }
}

func main() {
    jobs := make(chan int, 5)

    var wg sync.WaitGroup

    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }

    close(jobs)

    wg.Wait()
}
