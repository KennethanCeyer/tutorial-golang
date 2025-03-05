package main

import "fmt"

// Sum 함수는 제네릭 타입 T의 슬라이스와 합계에 대한 제약 조건을 정의
// T는 정수 또는 부동 소수점 숫자이어야 한다
func Sum[T int | float64](numbers []T) T {
    var total T
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    // 정수와 부동 소수점 숫자의 슬라이스에 대해 Sum 함수를 호출
    fmt.Println(Sum([]int{1, 2, 3, 4, 5}))         // 정수
    fmt.Println(Sum([]float64{1.1, 2.2, 3.3}))     // 부동 소수점 숫자
}
