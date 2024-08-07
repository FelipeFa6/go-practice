package main

import (
    "fmt"
)

func main() {
    // Expected output: 3
    arr1 := []int{1, 3, 5, 7, 9, 11}
    t1 := 7
    fmt.Println(linear(arr1, t1))

    // Expected output: -1
    arr2 := []int{2, 4, 6, 8, 10}
    t2 := 5
    fmt.Println(linear(arr2, t2))
}


func linear(arr []int, t int) (int, error) {
    for k, v := range arr {
        if(v == t) {
            return k, nil
        }   
    }

    return -1, fmt.Errorf("not found")
}
