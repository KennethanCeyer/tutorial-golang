package main

import "fmt"

func main() {
	// age 변수를 정의하고 30으로 초기화
	var age = 30

	// age2 변수를 정의하고 50으로 초기화
	// age2는 중복되지 않은 변수 이름이기 때문에 에러가 발생하지 않는다.
	var age2 = 50

	// 정의된 변수를 출력
	fmt.Println(age, age2)
}
