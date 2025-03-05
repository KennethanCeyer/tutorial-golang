package main

import "fmt"

func main() {
	operand := 5

	// 기존의 방식의 덧셈 연산
	sum := 10

	// 자기 자신의 항을 적어서 operand(5)와 더해
	// 자기 자신 sum 변수에 값을 대입한다.
	sum = sum + operand
	fmt.Println(sum) // 15


	// 신택틱 슈거 방식의 덧셈 연산
	sumSyntacticSugar := 10

	// 자기 자신에 덧셈을 할 경우 아래와 같이
	// 생략 문법을 사용할 수 있다.
	sumSyntacticSugar += operand
	fmt.Println(sumSyntacticSugar) // 15

	difference := 10
	difference -= 5
	fmt.Println(difference) // 5

	product := 5
	product *= 4
	fmt.Println(product) // 20

	quotient := 20
	quotient /= 5
	fmt.Println(quotient) // 4

	remainder := 10
	remainder %= 3
	fmt.Println(remainder) // 1
}
