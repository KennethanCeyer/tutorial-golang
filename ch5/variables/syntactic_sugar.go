package main

import "fmt"

func main() {
    // var 키워드 이후에 괄호로 묶어 변수를 정의할 수 있다.
    var (
        name = "Gil Dong, Hong"
        age = 30
        address = "Seoul, Republic of Korea"
    )

    fmt.Println(name, age, address)
}
