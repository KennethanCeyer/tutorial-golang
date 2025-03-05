package main

import "fmt"

func main() {
    // :=는 오직 변수를 초기화할 때만 사용할 수 있다.
    message := "hello dear readers"

    // 변수는 마찬가지로 다음과 같이 업데이트할 수 있다.
    // 문자열에 + 연산자를 이용하면 문자열을 이어붙일 수 있다.
    message += " nice to see you."

    fmt.Println(message)
}
