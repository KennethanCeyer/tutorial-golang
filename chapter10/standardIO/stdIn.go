package main

import "fmt"


func main() {
    var name string
    var age int


    // 사용자로부터 입력 받기
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    
    fmt.Print("Enter your age: ")
    fmt.Scanf("%d", &age)


    // 입력 받은 데이터 출력하기
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
