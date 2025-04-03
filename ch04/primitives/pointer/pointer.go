package main

import "fmt"

func main() {
	score := 100

	fmt.Printf("score 지역변수에 저장된 값: %d\n", score)
	fmt.Printf("score 지역변수의 주솟값: %d\n", &score)
	fmt.Printf("score 지역변수의 주솟값 (%%p): %p\n", &score)
}
