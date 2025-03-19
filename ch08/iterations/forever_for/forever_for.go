package main

import "fmt"

func main() {
	// 조건식에 해당하는 부분을 무조건 참(true)이 되도록 바꾸었다.
	for i := 0; true; i++ {
		n := i + 1
		fmt.Printf("%d번째 반복문: i=%d\n", n, i)
	}
}
