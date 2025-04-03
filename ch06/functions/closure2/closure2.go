package main

import "fmt"

// 바깥쪽 함수
func funcFactory(x int) func(int) int {
	// 안쪽 함수
	return func(y int) int {
		return x + y
	}
}

func main() {
	two := funcFactory(2)
	three := funcFactory(3)
	fmt.Printf("two 함수로 호출하면 %d\n", two(10))
	fmt.Printf("three 함수로 호출하면 %d", three(10))
}
