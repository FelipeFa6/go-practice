package main

import "fmt"

func main() {
    arr := []int{1,2,3}

    arr = arr[:5]
    printSlice(arr)
}
func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

