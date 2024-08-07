package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func get_weather(city string, wg *sync.WaitGroup) {
    defer wg.Done()

    var url string = "https://wttr.in/" + city + "?format=1"

    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error:%v\n", err)
        os.Exit(1)
    }

    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error:%v\n", err)
        os.Exit(1)
    }

    fmt.Println(string(body))
}

func main() {
    cities := []string{"London", "New_York", "Tokyo", "Paris", "Sydney"}

    var wg sync.WaitGroup

    for _, city := range cities {
        wg.Add(1)
        go get_weather(city, &wg)
    }

    wg.Wait()
}
