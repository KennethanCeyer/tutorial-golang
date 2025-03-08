package main

import "fmt"

func main() {
	// `12`는 10진수 정수 타입이므로
	// 암묵적으로 hour 변수는 int형으로 정의된다.
	var hour = 12

	// int형 hour 변수에 string형 값은 지정할 수 없다.
	hour = "six" // 오류 발생!

	fmt.Println(hour)
}
