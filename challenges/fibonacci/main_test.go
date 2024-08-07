package main

import (
    "testing"
    "fmt"
)

type test_t struct {
    n        int
    expected int
}

var tests = []test_t {
    {0, 0},
    {1, 1},
    {2, 1},
    {3, 2},
    {4, 3},
    {5, 5},
    {6, 8},
    {7, 13},
    {8, 21},
    {9, 34},
}

func TestGoodFibonacci(t *testing.T) {
    for _, test := range tests {
        t.Run(fmt.Sprintf("n=%d", test.n), func(t *testing.T) {
            if result := fibonacci(test.n); result != test.expected {
                t.Errorf("fibonacci(%d) = %d; expected %d", test.n, result, test.expected)
            }
        })
    }
}

