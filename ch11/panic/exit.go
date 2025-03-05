package main

import (
	"fmt"
	"os"
)


func main() {
    if err := doSomething(); err != nil {
        fmt.Println("에러 발생:", err)
        os.Exit(1) // 비정상적인 종료 코드 1을 설정
    }
    os.Exit(0) // 정상 종료 코드 0을 설정
}


func doSomething() error {
    // 에러를 발생시키는 가상의 함수
    return fmt.Errorf("문제가 발생했습니다.")
}
