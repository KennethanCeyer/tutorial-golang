package main

import "fmt"

func main() {
	// `12`는 10진수 정수 타입이므로
	// 암묵적으로 hour 변수는 int 타입으로 정의된다.
	var hour = 12

	// int 타입의 hour 변수에 string 타입의 값은 지정할 수 없게된다.
	hour = "six" // (에러발생!)

	fmt.Println(hour)
}
