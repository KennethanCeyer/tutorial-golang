package main

import "fmt"

func main() {
    // const는 상수를 지정할 때 사용하는 키워드.
    const pi = 3.1415926535898932
    const gravityConstant = 9.79641227572363

    // 프로그래머가 상수에 값을 새로 지정하려고 하고 있다.
    // 이 과정에서 에러가 발생할 것이다.
    pi = 4.231

    // fmt.Printf에 %v 서식 지정자에 상수가 대입되어 출력된다.
    fmt.Printf("파이 값은 %v입니다.\n", pi)
    fmt.Printf("중력 상수 G 값은 %v입니다.\n", gravityConstant)
}
