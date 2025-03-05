package main

import (
	"fmt"
	"sort"
	"time"
)

// 연산이 요구되는 무거운 작업
func compute(data []int) int {
	result := 0
	for _, value := range data {
		if value%2 == 0 {
			result += value * 2
		} else {
			result -= value * 2
		}
	}
	return result
}

func main() {
	data := make([]int, 1000000000)
	for i := range data {
		data[i] = i % 100
	}

	sort.Ints(data)

	start := time.Now()
	result := compute(data)
	elapsed := time.Since(start)

	fmt.Printf("결과: %d\n", result)
	fmt.Printf("실행 시간: %s\n", elapsed)
}
