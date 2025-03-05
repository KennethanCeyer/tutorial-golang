package main

import "fmt"

// value는 int형 변수의 주소값을 전달받는 타입
// 이런 형태를 포인터 변수라고 부른다
func updateValue(value *int) {
	// value 포인터를 참조해서, 주소의 위치에 할당된 값을 100으로 변경
	*value = 100
}

func main() {
	myValue := 50
	fmt.Printf("myValue 변경 전: %d\n", myValue)

	// myValue 변수의 주소 값을 전달
	updateValue(&myValue)
	fmt.Printf("myValue 변경 후: %d\n", myValue)
}
