package main

import "fmt"

func main() {
    // const는 상수를 지정할 때 사용하는 키워드.
    const pi = 3.1415926535898932
    const gravityConstant = 9.79641227572363

    // fmt.Printf에 %v 서식 지정자에 상수가 대입되어 출력된다.
    fmt.Printf("파이 값은 %v입니다.\n", pi)
    fmt.Printf("중력 상수 G 값은 %v입니다.\n", gravityConstant)
}
