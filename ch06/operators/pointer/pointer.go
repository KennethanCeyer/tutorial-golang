package main

import "fmt"

func main() {
    // 변수 선언 및 초기화
    var x int = 10

    // 참조 연산자
    var p *int = &x
    fmt.Println(p)  // 메모리 주소 출력

    // 역참조 연산자
    fmt.Println(*p) // 변수 x의 값 출력

    // 포인터를 통해 값 변경
    *p = 20
    fmt.Println(x)  // 20 출력
}
