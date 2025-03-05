package main

import "fmt"

func main() {
    // const 키워드를 var로 수정해도 동작에 문제는 발생하지 않는다.
    var pi = 3.1415926535898932
    var gravityConstant = 9.79641227572363

    // fmt.Printf에 %v 서식 지정자에 이번에는 변수가 대입되어 출력된다.
    fmt.Printf("파이 값은 %v입니다.\n", pi)
    fmt.Printf("중력 상수 G 값은 %v입니다.\n", gravityConstant)
}
