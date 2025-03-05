package main

import "fmt"

func AddPositiveNumbers(a, b int) int {
    if a < 0 || b < 0 {
        fmt.Println("Error: Negative numbers are not allowed")
        return 0
    }
    return a + b
}

