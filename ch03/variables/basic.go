package main

import "fmt"

func main() {
    // 상자를 하나 만들어 “Hello Go!”라는 문자열의 주소를 x라고 이름을 짓는다.
    var x = "Hello Go!"
    // 앞으로는 x라는 이름을 이용해 “Hello Go!”를 출력할 수 있다.
    fmt.Println(x)
    fmt.Println(x)
    fmt.Println(x)
}
