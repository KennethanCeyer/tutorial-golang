package main

import "fmt"

func main() {
    // 이제 변수 값을 지정하여 초기화하지 않고 생략할 수 있다.
    var count int

    // count에는 정수 형태의 숫자를 지정할 수 있다.
    count = 3

    // count에 문자열(string)의 값을 지정하면 에러가 발생하게 된다.
    count = "Hello World!" // (에러발생!)

    fmt.Println(count)
}
