package main

import "fmt"

func main() {
    fmt.Println("fmt.Println 출력")
    // fmt.Println() 함수는 문자열을 출력하고 개행한다.
    fmt.Println("apple")
    fmt.Println("banana")
    fmt.Println("========\n")

    fmt.Println("fmt.Printf 출력")
    // fmt.Printf() 함수는 자동으로 줄 넘김을 하지 않는다.
    fmt.Printf("apple")
    fmt.Printf("banana")
    fmt.Println("\n========\n")

    fmt.Println("fmt.Printf + 개행 문자 출력")
    // fmt.Printf() 함수는 개행 문자(\n)를 사용해야 한다.
    fmt.Printf("apple\n")
    fmt.Printf("banana\n")
    fmt.Println("========\n")

    // fmt.Printf는 %v 서식 지정자를 인식하고 대입한다.
    fmt.Printf("fmt.Printf: %v\n", "hello world")

    // fmt.Println은 %v를 그대로 출력하게 된다.
    fmt.Println("fmt.Println: %v", "hello world")
}
