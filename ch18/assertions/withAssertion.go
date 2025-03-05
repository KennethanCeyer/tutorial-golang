package main

func AddPositiveNumbersWithAssertion(a, b int) int {
    if a < 0 || b < 0 {
        panic("Negative numbers are not allowed")
    }
    return a + b
}

