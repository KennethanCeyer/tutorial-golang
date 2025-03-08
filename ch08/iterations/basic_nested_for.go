package main

import "fmt"

// 특정 단을 계산해서 출력해주는 함수 (EX: 2단, 4단)
func displayMultiplicationTableAt(tableIndex int) {
	for i := 1; i <= 9; i++ {
		fmt.Printf("%dx%d = %d\n", tableIndex, i, tableIndex*i)
	}
}

// 2단부터 9단까지 반복하며 구구단 출력
func displayMultiplicationTableAll() {
	for i := 2; i <= 9; i++ {
		displayMultiplicationTableAt(i)
		fmt.Print("\n")
	}
}

func main() {
	// 사용자로부터 받은 정수를 저장하는 변수
	var tableStartIndex int

	// 초기 프로그램 시작 문구
	fmt.Println("=================")
	fmt.Println("구구단 프로그램")
	fmt.Println("=================")

	// 사용자에게 몇단을 출력할지 물어보는 메시지
	fmt.Print("몇 번째 구구단을 출력할까요? ")
	fmt.Print("(0: 프로그램 종료): ")

	// 사용자로부터 정수를 입력받음
	fmt.Scanf("%d", &tableStartIndex)

	// 입력받은 정수가 0일경우 프로그램 종료
	if tableStartIndex == 0 {
		fmt.Println("프로그램을 종료합니다.")
		return
	}

	if tableStartIndex >= 2 && tableStartIndex <= 9 {
		// 입력받은 정수가 2~9 값이면 특정 단 출력
		displayMultiplicationTableAt(tableStartIndex)
	} else {
		// 입력받은 정수가 2~9 이외의 값이면 전체 구구단 출력
		displayMultiplicationTableAll()
	}
}
