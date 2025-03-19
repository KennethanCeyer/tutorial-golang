package main

import "fmt"

func main() {
	// const는 상수를 지정할 때 사용하는 키워드이다.
	const pi = 3.1415926535898932
	const gravityConstant = 9.79641227572363

	// 상수에 값을 새로 지정한다.
	pi = 4.231 // 상수는 초기화 후 값을 변경할 수 없으므로 오류 발생!

	fmt.Printf("파이 값은 %v입니다.\n", pi)
	fmt.Printf("중력 상수 G 값은 %v입니다.\n", gravityConstant)
}
