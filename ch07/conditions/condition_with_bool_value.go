package main

import "fmt"

func checkConditionWithBool() {
	// 무조건 참인 경우
	if true {
		fmt.Println("Checked true!")
	} else {
		// 3번째 행에 true를 조건 값으로 사용하였으므로,
		// 이 분기는 실행되지 않음.
		fmt.Println("Check false!")
	}
}

func main() {
	checkConditionWithBool()
}
