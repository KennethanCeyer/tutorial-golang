package main


import (
    "testing"
)


func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sum(1, 2)
    }
}


func sum(a, b int) int {
    return a + b
}

