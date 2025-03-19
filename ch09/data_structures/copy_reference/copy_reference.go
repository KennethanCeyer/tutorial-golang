package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	reference := original

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
