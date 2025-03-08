package main

import "fmt"

func main() {
    // age 변수를 정의하고 30으로 초기화
    age := 30
    
    // age 변수를 정의하고 50으로 초기화
    age := 50 // age 변수가 이미 정의되어있음, (에러발생!)

    // 정의된 변수를 출력
    fmt.Println(age)
}

