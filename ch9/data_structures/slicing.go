package main

import "fmt"

func main() {
	// 6 길이를 가지는 리터럴 슬라이스
	original := []int{1, 2, 3, 4, 5, 6}

	// 두번째 요소부터 네번째 요소까지의 값을 슬라이싱
	slicing := original[1:4]

	fmt.Println("================================================")
	fmt.Println("[original]")
	fmt.Printf("len: %d, cap: %d\n", len(original), cap(original))
	fmt.Printf("value: %v\n", original)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[slicing]")
	fmt.Printf("len: %d, cap: %d\n", len(slicing), cap(slicing))
	fmt.Printf("value: %v\n", slicing)
	fmt.Println("================================================")
}
