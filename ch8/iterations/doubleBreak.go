package main

import "fmt"

func main() {
for i := 0; i < 5; i++ {
	for j := 0; j < 3; j++ {
		fmt.Printf("i=%d, j=%d\n", i, j)
		if j == 1 {
			// 7번째 행의 반복문 탈출
			break

			// 6번째 행의 반복문 탈출
			break
		}
	}
}
}
