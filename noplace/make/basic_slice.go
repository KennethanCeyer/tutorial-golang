package main

import "fmt"

func main() {
	// nil 슬라이스 정의
	var nilSlice []int

	// 빈 슬라이스 정의
	emptySlice := make([]int, 0)

	// 8 길이를 가진 bool 리터럴 슬라이스 정의
	literalSlice := []bool{false, false, false, false}

	// 2 길이를 가진 byte 슬라이스 정의
	byteSlice := make([]byte, 2)

	// 4 길이를 가지고,
	// 9 용량을 가진 string 슬라이스 정의
	stirngSlice := make([]string, 4, 9)

	fmt.Println("================================================")
	fmt.Println("[nilSlice]")
	fmt.Printf("len: %d, cap %d\n", len(nilSlice), cap(nilSlice))
	fmt.Printf("is nil: %t\n", nilSlice == nil)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[emptySlice]")
	fmt.Printf("len: %d, cap %d\n", len(emptySlice), cap(emptySlice))
	fmt.Printf("is nil: %t\n", emptySlice == nil)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[literalSlice]")
	fmt.Printf("len: %d, cap %d\n", len(literalSlice), cap(literalSlice))
	fmt.Printf("is nil: %t\n", literalSlice == nil)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[byteSlice]")
	fmt.Printf("len: %d, cap %d\n", len(byteSlice), cap(byteSlice))
	fmt.Printf("is nil: %t\n", byteSlice == nil)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[stirngSlice]")
	fmt.Printf("len: %d, cap %d\n", len(stirngSlice), cap(stirngSlice))
	fmt.Printf("is nil: %t\n", stirngSlice == nil)
	fmt.Println("================================================")
}
