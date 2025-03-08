package main

import (
	"fmt"
)


func main() {
    var name string


    // 사용자에게 이름을 입력받는다
    fmt.Print("이름을 입력해주세요: ")
    fmt.Scanf("%s", &name)


    // 입력받은 이름을 출력한다
    fmt.Printf("안녕하세요, %s님!\n", name)
}
