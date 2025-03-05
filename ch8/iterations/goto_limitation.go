package main

import "fmt"

func main() {
	fmt.Println("10번째 행로 이동")
	goto inside_loop

	for i := 0; i < 5; i++ {
	inside_loop:
		fmt.Printf("%d번째 반복문\n", i)
	}
}
