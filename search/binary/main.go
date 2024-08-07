package main

import(
    "fmt"
)

func main() {
    var a = []int {1, 2, 3, 4, 5, 6}
    res := binary_search(a, 2, 0)

    if(res == -1) {
        fmt.Println("Not found!")
    } else {
        fmt.Printf("Found at index: %d\n", res)
    }
}

func binary_search(array []int, target int, offset int) int {
    // target not found
    if(len(array) == 0) {
        return -1
    }

    var middle int = len(array)/2
    var c int = array[middle]

    if (c == target) {
        return offset + middle
    }

    if(c > target) {
        return binary_search(array[0:middle], target, offset)
    } else {
        return binary_search(array[middle+1:], target, offset+middle+1)
    }
}
