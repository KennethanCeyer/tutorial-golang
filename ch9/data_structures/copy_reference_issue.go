package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	reference := original

	// original 변수의 0번 인덱스 요소 값 변경
	original[0] = 5

	fmt.Println("================================================")
	fmt.Println("[original]")
	fmt.Printf("len: %d, cap: %d\n", len(original), cap(original))
	fmt.Printf("value: %v\n", original)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[reference]")
	fmt.Printf("len: %d, cap: %d\n", len(reference), cap(reference))
	fmt.Printf("value: %v\n", reference)
	fmt.Println("================================================")
}
