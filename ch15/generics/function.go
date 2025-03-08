package main

import "fmt"

// 제네릭 함수 정의
func Swap[T any](a, b T) (T, T) {
    return b, a
}


func main() {
    // 정수 타입의 값 교환
    x, y := 1, 2
    x, y = Swap(x, y)
    fmt.Println("Swapped:", x, y)  // 출력: Swapped: 2 1


    // 문자열 타입의 값 교환
    str1, str2 := "hello", "world"
    str1, str2 = Swap(str1, str2)
    fmt.Println("Swapped:", str1, str2)  // 출력: Swapped: world hello
}
