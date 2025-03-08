package main

import "fmt"

func main() {
  const (
    name = "홍길동"
    age = 31
  )
  const pi, gravityConstant = 3.1415926535898932, 9.79641227572363

  fmt.Printf("안녕하세요 저는 %v이고 나이는 %d입니다.\n", name, age)
  fmt.Printf("파이 값은 %v입니다.\n", pi)
  fmt.Printf("중력 상수 G 값은 %v입니다.\n", gravityConstant)
}
