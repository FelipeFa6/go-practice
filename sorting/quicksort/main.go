package main

import "fmt"

func main() {
    t1 := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println(quicksort(t1, 0, len(t1 -1)))
}

func quicksort(arr []int, start int, end int) []int{
    if (end <= start) {
        return arr
    }

    pivot := partition(arr, start, len(arr) - 1)
    quicksort(arr, 0, 2)
    return arr
}

func partition(arr []int, start int, end int) int{
    i := 1
    return i
}
