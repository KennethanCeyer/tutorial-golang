package main

import "fmt"

// a, b 이름의 정수형(int) 자료형 인자를 받고,
// 2개의 정수형 반환을 갖는 함수 정의.
func addAndSub(a int, b int) (int, int) {
	add := a + b
	sub := a - b

	// add, sub 변수 총 2개의 값을 반환.
	return add, sub
}

func main() {
	// 정수형 자료형의 함수의 입력 값을 변수로 정의함.
	inputA := 2
	inputB := 3

	// addAndSub 함수에 inputA, inputB 인자를 이용해 호출하고 반환 값을
	// 각각 addResult, subResult 변수에 정의.
	addResult, subResult := addAndSub(inputA, inputB)

	// addResult, subResult에 저장된 값을 출력 함.
	fmt.Printf("%v + %v = %v\n", inputA, inputB, addResult)
	fmt.Printf("%v - %v = %v\n", inputA, inputB, subResult)
}
