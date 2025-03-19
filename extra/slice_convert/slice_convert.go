package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}

	// 변환한 값을 저장하기 위해
	// 기존 슬라이스와 동일한 크기의 슬라이스 할당
	doubled := make([]int, len(original))

	// 변환하는 코드 작성
	for i, value := range original {
		doubled[i] = value * 2
	}

	fmt.Printf("기존 슬라이스: %v\n", original)
	fmt.Printf("2배로 곱해진 슬라이스: %v\n", doubled)
}
