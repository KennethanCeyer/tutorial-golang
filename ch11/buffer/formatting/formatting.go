package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("이름과 나이를 입력하세요 (예: 홍길동 30): ")
	fmt.Scanf("%s %d\n", &name, &age) // 서식 문자열 사용
	fmt.Printf("입력된 이름: %s, 나이: %d\n", name, age)
}
