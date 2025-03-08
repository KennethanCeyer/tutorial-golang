package main

import "fmt"

func main() {
	conditionValue := 10

if conditionValue > 5 {
	fmt.Println("OK!")
} else {
	fmt.Println("Not good!: 모든 조건에 해당되지 않습니다!")
} else {
	fmt.Println("조건 값이 범위를 벗어났습니다!")
}
}
