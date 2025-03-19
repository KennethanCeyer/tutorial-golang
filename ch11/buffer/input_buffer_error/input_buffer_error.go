package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("첫 번째 숫자를 입력하세요: ")
	fmt.Scan(&a) // 개행 문자가 입력 버퍼에 남음

	fmt.Print("두 번째 숫자를 입력하세요: ")
	fmt.Scanln(&b) // 이전 입력의 개행 문자로 인해 입력 처리 종료

	fmt.Printf("입력된 값: %d, %d\n", a, b)
}
