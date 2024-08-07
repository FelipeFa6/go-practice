package main

import (
    "fmt"
)

// 0, 1, 1, 2, 3, 5, 8, 13
func fibonacci(n int) int {
    if (n <= 1) {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    r := fibonacci(6)
    fmt.Println(r)
}
