package main

import "fmt"

func main() {
	// 모든 산술 연산자는 이항 연산자 형태입니다.
	//       좌측항   연산자    우측항
	format :=  1      +       2
	fmt.Println(format) // 3

	// 좌측 항과 우측 항을 더함
	sum := 5 + 6
	fmt.Println(sum) // 11

	// 좌측 항에 우측 항을 뺌
	difference := 5 - 6
	fmt.Println(difference) // -1

	// 좌측 항과 우측 항을 곱함
	product := 5 * 6
	fmt.Println(product) // 30

	// 좌측 항에 우측 항을 나눔
	quotient := 6 / 2
	fmt.Println(quotient) // 3

	// 좌측 항에 우측 항을 나눠 나머지를 구함
	remainder := 10 % 3
	fmt.Println(remainder) // 1
}
