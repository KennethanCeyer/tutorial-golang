package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	reference := make([]int, len(original))

	// original에 있는 요소 값을
	// reference에 깊은 복사(deep copy)
	copy(reference, original)

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
