package main

import "fmt"

func main() {
	var nilValue []string

	// nilValue에 값을 할당하기 전에 값이 nil인지 체크
	fmt.Printf("Is Nil?: %t\n", nilValue == nil)

	nilValue = []string{"hello"}

	// nilValue에 값을 할당하고 값이 nil인지 체크
	fmt.Printf("Is Nil?: %t\n", nilValue == nil)
}
