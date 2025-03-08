package main

import "fmt"


func main() {
    name := "Alice"
    age := 30


    // 기본적인 출력
    fmt.Print("Name: ")
    fmt.Print(name)
    fmt.Print("\n")


    // 자동 줄바꿈
    fmt.Println("Age:", age)


    // 포맷된 출력
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
