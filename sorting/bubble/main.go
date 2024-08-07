package main

import(
    "fmt"
)

func main() {
    t1 := []int{64, 34, 25, 12, 22, 11, 90}

    fmt.Println(bubble_sort(t1))
}

func bubble_sort(arr []int) []int {
    for i := 0; i < len(arr) -1; i++ {
        for j := 0; j < len(arr) - i - 1; j++ {
            if(arr[j] > arr[j+1]) {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    return arr
}
