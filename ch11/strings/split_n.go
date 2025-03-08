package main

import (
	"fmt"
	"strings"
)

func main() {
	original := "hello world"
	separator := " "
	result := strings.SplitN(original, separator, 2)

	fmt.Printf("기존 문자열: '%s'\n", original)
	fmt.Printf("분리 기준 문자열: '%s'\n", separator)

	fmt.Println("================================================")
	fmt.Printf("분리된 문자열 슬라이스: %v\n", result)
}
