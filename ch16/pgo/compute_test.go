package main

import (
	"testing"
)

func BenchmarkCompute(b *testing.B) {
	data := make([]int, 1000000000)
	for i := range data {
		data[i] = i % 100 
	}

	for i := 0; i < b.N; i++ {
		_ = compute(data)
	}
}
