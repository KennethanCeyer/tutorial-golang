package main

import "fmt"

// 제네릭 타입 정의
type Pair[T any] struct {
    First  T
    Second T
}


// 제네릭 함수 정의
func PrintPair[T any](p Pair[T]) {
    fmt.Printf("First: %v, Second: %v\n", p.First, p.Second)
}


func main() {
    // 정수 타입을 사용하는 Pair
    intPair := Pair[int]{1, 2}
    PrintPair(intPair)


    // 문자열 타입을 사용하는 Pair
    strPair := Pair[string]{"hello", "world"}
    PrintPair(strPair)
}
