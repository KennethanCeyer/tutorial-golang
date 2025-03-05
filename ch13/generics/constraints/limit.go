package main

import "fmt"

// 제약 조건 없는 제네릭 함수 정의
func PrintLength[T any](value T) {
    fmt.Println(len(value))  // 이 코드는 T가 len 함수를 지원해야 한다.
}

func main() {
    PrintLength("hello") // 정상 작동
    PrintLength([]int{1, 2, 3}) // 정상 작동
    // PrintLength(42) // 컴파일 오류: int 타입은 len 함수를 지원하지 않는다.
}
