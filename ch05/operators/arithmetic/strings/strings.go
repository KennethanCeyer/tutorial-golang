package main

import "fmt"

func main() {
	text1 := "hello"
	text2 := "world"

	// string 값을 + 연산자를 이용하여 연산하면
	// 문자열을 연결 시키는 형식으로 동작한다.
	message := text1 + " " + text2
	fmt.Println(message)
}
