package main

import "fmt"

func main() {
	// nil 슬라이스 sliceValue 변수 선언
	var sliceValue []int

	// sliceValue 변수에 4가지 요소를 추가
	sliceValue = append(sliceValue, 1, 2, 3, 4)

	// sliceValue 변수 값 출력
	fmt.Printf("sliceValue len: %d, cap: %d, value: %v\n",
		len(sliceValue), cap(sliceValue), sliceValue)
}
