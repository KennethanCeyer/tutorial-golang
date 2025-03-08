package main

import "fmt"

func main() {
    // int 타입의 포인터 변수 선언
    var ptr *int

    // int 타입 변수 선언 및 초기화
    var val int = 42

    // ptr에 val 변수의 주소 할당
    ptr = &val

    // ptr이 가리키는 주소의 값 출력 (역참조)
    fmt.Printf("ptr이 가리키는 값: %d\n", *ptr)

    // ptr이 가리키는 주소에 새로운 값 할당
    *ptr = 100

    // val 변수의 값 출력 (ptr을 통해 수정된 값)
    fmt.Printf("val의 새로운 값: %d\n", val)
}
