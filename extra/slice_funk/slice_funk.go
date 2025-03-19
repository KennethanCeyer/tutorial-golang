package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

func main() {
	original := []int{1, 2, 3, 4, 5}

	// funk.Map의 고차 함수를 이용해 변환
	doubled := funk.Map(original, func(value int) int {
		return value * 2
	}).([]int)

	fmt.Printf("기존 슬라이스: %v\n", original)
	fmt.Printf("2배로 곱해진 슬라이스: %v\n", doubled)
}
