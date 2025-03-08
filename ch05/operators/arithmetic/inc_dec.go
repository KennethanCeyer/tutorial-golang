package main

import "fmt"

func main() {
	// 기존 방식
	increase := 1
	increase += 1
	fmt.Println(increase) // 2

	// 증가 연산자
	increaseSyntacticSugar := 1
	increaseSyntacticSugar++
	fmt.Println(increaseSyntacticSugar) // 2

	// 기존 방식
	decrease := 1
	decrease -= 1
	fmt.Println(decrease) // 0

	// 감소 연산자
	decreaseSyntacticSugar := 1
	decreaseSyntacticSugar--
	fmt.Println(decreaseSyntacticSugar) // 0
}
