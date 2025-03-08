package main

import "fmt"

// Stringer 인터페이스 정의
type Stringer interface {
    String() string
}

// Error 인터페이스 정의
type Error interface {
    Error() string
}

// 제네릭 타입 제약 정의
type StringerError interface {
    Stringer
    Error
}

// 제네릭 함수 정의
func PrintInfo[T StringerError](value T) {
    fmt.Println(value.String())
    fmt.Println(value.Error())
}

// Stringer 및 Error 인터페이스를 구현하는 타입 정의
type MyError struct {
    Message string
}

func (e MyError) String() string {
    return "String() => " + e.Message
}

func (e MyError) Error() string {
    return "Error() => " + e.Message
}

func main() {
    // MyError 타입을 사용하는 제네릭 함수
    e := MyError{Message: "Something went wrong"}
    PrintInfo(e)
}
