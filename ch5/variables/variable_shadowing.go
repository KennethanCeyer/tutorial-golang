package main

import "fmt"

// 전역변수 name 선언
var (
	name = "손님"
)

func display() {
	fmt.Printf("<display> 안녕하세요! 현재 이름은: %s 이십니다\n", name)

	// 지역변수 name 선언
	// 전역변수와 동일한 이름으로 정의되었다
	var name string
	fmt.Print("혹시 이름을 알려주실 수 있으신가요?: ")
	fmt.Scanf("%s", &name)

	// display 함수에서 name 변수 이름에 접근하여 값 출력
	fmt.Println("================================================")
	fmt.Printf("<display> 안녕하세요! '%s'님 현재 이름은: %s 이십니다\n",
		name, name)
}

func main() {
	display()

	// main 함수에서 name 변수 이름에 접근하여 값 출력
	fmt.Println("================================================")
	fmt.Printf("<main> 안녕하세요! '%s'님 현재 이름은: %s 이십니다\n",
		name, name)
}
