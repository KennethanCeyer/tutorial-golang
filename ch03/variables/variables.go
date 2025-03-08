package main

import "fmt"

func main() {
    // 상자를 하나 만들어 1을 담고 total이라고 이름을 짓는다.
    var total = 1

    // total이라는 상자에 3을 더한다.
    total = total + 3

    // total이라는 상자에 total만큼 곱한다.
    total = total * total

    // total 변수 값을 출력한다.
    fmt.Println(total)
}
