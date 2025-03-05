package main

import (
	"fmt"
	"unsafe"
)

func main() {
    var charA byte = 'a'
    var stringA string = "a"
    var stringHello string = "hello world!"

    fmt.Println("format")
    fmt.Println("========")

    // byte는 일반 출력에서는 숫자 형식으로 출력된다.
    fmt.Printf("byte charA: %v(number)\n", charA)

    // 서식 지정자 %c를 통해 문자 형태로 출력이 가능하다.
    fmt.Printf("byte charA: %c(char)\n", charA)

    // 문자를 문자열 형식으로 출력하면 올바르게 출력되지 않는다.
    fmt.Printf("byte charA: %s(string)\n\n", charA)

    // 서식 지정자 %s를 통해 문자열 형태로 출력이 가능하다.
    fmt.Printf("string stringA: %s(string)\n", stringA)

    // 문자열을 문자 형식으로 출력하면 올바르게 출력되지 않는다.
    fmt.Printf("string stringA: %c(char)\n", stringA)

    // 문자열은 일반적으로 둘 이상의 문자를 출력할 때 사용한다.
    fmt.Printf("string stringHello: %v\n\n", stringHello)

    fmt.Println("size")
    fmt.Println("========")

    // 문자는 바이트 자료형을 기반하기 때문에 1바이트이다.
    fmt.Printf("char size: %d\n", unsafe.Sizeof(charA))

    // 문자열은 16바이트를 차지하게 된다.
    fmt.Printf("string size: %d\n", unsafe.Sizeof(stringA)) 

    // 문자열의 길이와 관계 없이 16바이트를 차지한다.
    fmt.Printf("large string size: %d", unsafe.Sizeof(stringHello))
}
