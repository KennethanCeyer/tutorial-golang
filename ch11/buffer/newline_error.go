package main

import "fmt"

func main() {
   var name string
   var age int
   var message string

   fmt.Scanf("%s %d", &name, &age)
   fmt.Scanln(&message)   // 남은 개행 문자로 인해 입력 없이 종료
}
