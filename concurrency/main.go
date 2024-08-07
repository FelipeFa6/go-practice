package main

import (
	"fmt"
	"sync"
	"time"
)

func do_work(t time.Duration, s string, wg *sync.WaitGroup, ch chan<-string) {
    defer wg.Done()

    fmt.Println("wait...")
    time.Sleep(t)
    ch <- s
}

func main() {
    var wg sync.WaitGroup
    r_ch := make(chan string)

    wg.Add(4)
    go do_work(time.Second * 1, "Hello", &wg, r_ch)
    go do_work(time.Second * 3, ", ", &wg, r_ch)
    go do_work(time.Second * 1, "World", &wg, r_ch)
    go do_work(time.Second * 4, "World!!", &wg, r_ch)

    go func() {
        wg.Wait()
        close(r_ch)
    }()

    for ch := range r_ch {
        fmt.Println(ch)
    }

    fmt.Println("---Finished---")
}
