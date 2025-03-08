package main

import "fmt"

// Stringer와 Error 인터페이스를 모두 만족하는 타입을 제약 조건으로 하는 제네릭 함수
func PrintDetails[T fmt.Stringer & error](value T) {
    fmt.Println("String:", value.String())
    fmt.Println("Error:", value.Error())
}

// Stringer와 Error 인터페이스를 구현하는 타입
type MyError struct {
    Message string
}

func (e MyError) String() string {
    return e.Message
}

func (e MyError) Error() string {
    return e.Message
}

func main() {
    // MyError 타입을 사용하는 제네릭 함수 호출
    e := MyError{Message: "An error occurred"}
    PrintDetails(e)
}
