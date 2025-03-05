package main

import "fmt"

// 제네릭 타입 정의
type Pair[T any] struct {
    First  T
    Second T
}

func main() {
    // 정수 타입의 Pair
    intPair := Pair[int]{1, 2}
    fmt.Println("Integer Pair:", intPair)  // 출력: Integer Pair: {1 2}

    // 문자열 타입의 Pair
    strPair := Pair[string]{"hello", "world"}
    fmt.Println("String Pair:", strPair)  // 출력: String Pair: {hello world}
}
